package main

import (
    "fmt"
    "net/http"
    "code.google.com/p/go.net/html"
    "strings"
    "net/url"
    "io/ioutil"
    "time"
    "strconv"
    "io"
    "os"
    "archive/zip"
    "path/filepath"
    "sync"
)

type Tokens struct {
    ViewState, EventValidation string;
    FileNames [5]string; //max of four quarters and one new registration option a year
}

var savePath = "./filings/"
var tmpName = "tmp"
var fileExt = ".zip"
var link = "http://disclosures.house.gov/ld/LDDownload.aspx?KeepThis=true"

func Unzip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()
 
    for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
 
        path := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            f, err := os.OpenFile(
                path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer f.Close()
 
            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
    }
 
    return nil
}

func download(t Tokens, i int, wg *sync.WaitGroup) {
    //set POST data
    v := url.Values{}
	v.Set("__VIEWSTATE", t.ViewState)
	v.Set("__EVENTVALIDATION", t.EventValidation)
	v.Set("selFilesXML", t.FileNames[i])
	v.Set("btnDownloadXML", "Download")
    
    fmt.Println("Downloading " + t.FileNames[i] + " to " + savePath + tmpName + fileExt + "...")
    
    //pres, err := http.PostForm("http://vm-2.ansonl.koding.kd.io/php.php", v)
    pres, err := http.PostForm(link, v)
    
    robots, err := ioutil.ReadAll(pres.Body)
	pres.Body.Close()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(savePath + tmpName + strconv.Itoa(i) + fileExt, robots, 0644)
		if err != nil {
		panic(err)
	}
	
	fmt.Println("Downloaded " + t.FileNames[i] + " to " + savePath + tmpName + strconv.Itoa(i) + fileExt)
	
	//Unzipping
	fmt.Println("Unzipping " + savePath + tmpName + strconv.Itoa(i) + fileExt + " to " + savePath + "...")
	
	err = Unzip(savePath + tmpName + strconv.Itoa(i) + fileExt, savePath)
    if err != nil {
        panic(err)
    }
    
    fmt.Println("Unzipped " + savePath + tmpName + strconv.Itoa(i) + fileExt + " to " + savePath)
    
    //Waitgroup done
    wg.Done()
}

func scrape() {
    fmt.Println("Sending GET request to " + link + "...");
    
    res, err := http.Get(link)
    
    if (err != nil) {
        panic(err)
    }
    
    doc, err := html.Parse(res.Body)
    
    token := Tokens{}
    tokenFileNamesCount := 0;
    
    //documentation https://godoc.org/code.google.com/p/go.net/html#Attribute
    var f func(*html.Node)
    f = func(n *html.Node) {
        //'n' is a node representing ONE object on the page
        //fmt.Println(n.Attr)
        if n.Type == html.ElementNode {
            for _, a := range n.Attr {
                
                //fmt.Println(a)
                //fmt.Println(a.Namespace, a.Key, a.Val)
                if (len(a.Val) > 3 && strings.Contains(a.Val[0:4], strconv.Itoa(time.Now().Year())) && tokenFileNamesCount < 5) { //check first four characters of attr value to get year
                    token.FileNames[tokenFileNamesCount] = a.Val
                    tokenFileNamesCount++
                }
                if (a.Val == "__VIEWSTATE") {
                    //fmt.Println(n.Attr[2].Key, n.Attr[2].Val)
                    for _, b := range n.Attr { //loop through this same object again, looking for the attribute in the slice which has key of "value"
                        if (b.Key == "value") {
                            token.ViewState = b.Val
                        }
                    }
                }
                if (a.Val == "__EVENTVALIDATION") {
                    for _, b := range n.Attr { //loop through this same object again, looking for the attribute in the slice which has key of "value"
                        if (b.Key == "value") {
                            token.EventValidation = b.Val
                        }
                    }
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling { //advance node
            f(c)
        }
    }
    f(doc)
    
    //Download each file in a thread
    
    
    if _, err := os.Stat(savePath); err != nil {
        if os.IsNotExist(err) {
            err = os.Mkdir(savePath, 0777);
            if err != nil {
                panic(err)
            }
        } else {
            panic(err)
        }
    }
    
    fmt.Println(savePath + " directory made")
    
    var wg sync.WaitGroup
    for fileNumber, file := range token.FileNames {
        if file != "" {
            wg.Add(1)
            go download(token, fileNumber, &wg)            
        }
    }
    wg.Wait()
    fmt.Println("All files downloaded.")
}
