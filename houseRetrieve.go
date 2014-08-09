package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Tokens struct {
	ViewState, EventValidation string
	FileNames                  []string //max of four quarters and one new registration option a year
}

var savePathHouse = "./houseFilings/"
var tmpNameHouse = "tmp"
var fileExtHouse = ".zip"
var linkHouse = "http://disclosures.house.gov/ld/LDDownload.aspx?KeepThis=true"

func download(t Tokens, i int, file string, wg *sync.WaitGroup) {
	//set POST data
	v := url.Values{}
	v.Set("__VIEWSTATE", t.ViewState)
	v.Set("__EVENTVALIDATION", t.EventValidation)
	v.Set("selFilesXML", t.FileNames[i])
	v.Set("btnDownloadXML", "Download")

	fmt.Println("Downloading " + file + " to " + savePathHouse + tmpNameHouse + strconv.Itoa(i) + fileExtHouse + "...")

	//pres, err := http.PostForm("http://vm-2.ansonl.koding.kd.io/php.php", v)
	pres, err := http.PostForm(linkHouse, v)

	robots, err := ioutil.ReadAll(pres.Body)
	pres.Body.Close()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(savePathHouse+tmpNameHouse+strconv.Itoa(i)+fileExtHouse, robots, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Downloaded " + file + " to " + savePathHouse + tmpNameHouse + strconv.Itoa(i) + fileExtHouse)

	//Unzipping
	fmt.Println("Unzipping " + savePathHouse + tmpNameHouse + strconv.Itoa(i) + fileExtHouse + " to " + savePathHouse + "...")

	err = Unzip(savePathHouse+tmpNameHouse+strconv.Itoa(i)+fileExtHouse, savePathHouse)
	if err != nil {
		panic(err)
	}

	fmt.Println("Unzipped " + savePathHouse + tmpNameHouse + strconv.Itoa(i) + fileExtHouse + " to " + savePathHouse)

	//Waitgroup done
	wg.Done()
}

func Extend(slice []string, element string) []string {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]string, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func scrape() string {

	fmt.Println("Sending GET request to " + linkHouse + "...")

	res, err := http.Get(linkHouse)

	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(res.Body)

	token := Tokens{}
	token.FileNames = make([]string, 0)

	//documentation https://godoc.org/code.google.com/p/go.net/html#Attribute
	var f func(*html.Node)
	f = func(n *html.Node) {
		//'n' is a node representing ONE object on the page
		//fmt.Println(n.Attr)
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {

				//fmt.Println(a)
				//fmt.Println(a.Namespace, a.Key, a.Val)
				if len(a.Val) > 3 && strings.Contains(a.Val[0:2], strconv.Itoa(time.Now().Year())[0:2]) { //check first two characters of attr value to get year
					token.FileNames = Extend(token.FileNames, a.Val)
					//fmt.Println(a.Val)
				}
				if a.Val == "__VIEWSTATE" {
					//fmt.Println(n.Attr[2].Key, n.Attr[2].Val)
					for _, b := range n.Attr { //loop through this same object again, looking for the attribute in the slice which has key of "value"
						if b.Key == "value" {
							token.ViewState = b.Val
						}
					}
				}
				if a.Val == "__EVENTVALIDATION" {
					for _, b := range n.Attr { //loop through this same object again, looking for the attribute in the slice which has key of "value"
						if b.Key == "value" {
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

	if _, err := os.Stat(savePathHouse); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(savePathHouse, 0777)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	fmt.Println(savePathHouse + " records directory made")

	var wg sync.WaitGroup
	for fileNumber, file := range token.FileNames[len(token.FileNames)-6 : len(token.FileNames)] {
		if file != "" {
			wg.Add(1)
			go download(token, fileNumber, file, &wg)
		}
	}
	wg.Wait()
	fmt.Println("All House files downloaded.")

	return savePathHouse //return saved path
}

func downloadHouseData() string {
	return scrape()
}
