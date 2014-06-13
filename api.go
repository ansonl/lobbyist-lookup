package main

import (
	"encoding/xml"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"path/filepath"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Lobbyist struct {
	FirstName string `xml:"lobbyistFirstName"`
	LastName string `xml:"lobbyistLastName"`
}

type Registration struct {
	OrganizationName string `xml:"organizationName"`
	ClientName string `xml:"clientName"`
	SenateID string `xml:"senateID"`
	HouseID string `xml:"houseID"`
	//ReportYear string `xml:"reportYear"`
	//ReportType string `xml:"reportType"`
	//Lobbyist []Lobbyist `xml:"alis>ali_info>lobbyists>lobbyist"` //apparently house changed their xml format on 6/10??
	Lobbyist []Lobbyist `xml:"lobbyists>lobbyist"`
}

var rArray []Registration

//return a json formatted string for a day struct
func (reg Registration) JSONString() []byte {
	b, err := json.Marshal(reg)
	if (err != nil) {
		panic(err)
	}
	return b
}

func handler(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadFile("./pages/api.html")
    if (err != nil) {
		panic(err)
	}
    fmt.Fprintf(w, string(data))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

    	r.ParseForm()
    	fmt.Println(r.Form)
    
    	//bypass same origin policy
    	w.Header().Set("Access-Control-Allow-Origin", "*")
    
    	//setup return string
    	returnString := "{" + `"` + "array" + `"` + ":" + "[ ";
    
    	//firstName := r.Form["first"]
    	lastName := r.Form["last"]
    	companyName := r.Form["company"]
    
        limit := 10
        count := 0
    
    	for _, i := range rArray {
    	    if (count < limit) {
        		if (companyName != nil) {
        			for _, k := range companyName {
        				if (strings.Contains(strings.ToLower(i.OrganizationName), k)) {
        					if (lastName != nil) {
        						for _, j := range i.Lobbyist {
        							if (j.LastName != "") {
        								for _, l := range lastName {
        									if (strings.Contains(strings.ToLower(j.LastName), l)) {
        										returnString += string(i.JSONString()) + ","
        										count++
        									}
        								}
        							}
        						}
        					} else {
        						returnString += string(i.JSONString()) + ","
        						count++
        					}
        				}
        			}
        		} else {
        		    if (lastName != nil) {
            			for _, j := range i.Lobbyist {
            				if (j.LastName != "") {
            					for _, l := range lastName {
            						if (strings.Contains(strings.ToLower(j.LastName), l)) {
            							returnString += string(i.JSONString()) + ","
            							count++
            						}
            					}
            				}
            			}
        		    }
        		}
    	    }
    	}
    
    	returnString = returnString[:len(returnString) - 1]
    	returnString += "]" + "}"
    	fmt.Fprintf(w, returnString)
}

func server() {
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)
    
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil) 
    if err != nil {
      panic(err)
    }  

    fmt.Println("listening on port " + os.Getenv("PORT"))
}

func readDirectory(recordDir string) {
    files, err := ioutil.ReadDir("./" + recordDir + "/")
	if (err != nil) {
		panic(err)
	}

    fmt.Println("Reading " + strconv.Itoa(len(files)) + " files from " + recordDir + "...")

    rArray = nil
	rArray = make([]Registration, len(files))

	a := 0 //counter for number of files successfully read

	for _, f := range files {
            data, err := ioutil.ReadFile(recordDir + "/" + f.Name())
            if err != nil {
    			fmt.Println("error reading %v", err)
    			return
    		} else {
    	        if (strings.Contains(filepath.Ext(f.Name()), "xml")) {
    
    	        	//unmarshal data and put into struct array
    				err = xml.Unmarshal([]byte(data), &rArray[a])
    				if err != nil {
    					fmt.Println("error decoding %v: %v",f.Name(), err)
    					return
    				}
    
    				a++ //increment number of files successfully parsed
    			}
    		}

		if (a % 1000 == 0) {
		    fmt.Println(strconv.Itoa(a) + " files read");
		}
    }

    fmt.Println("Successfully read " , a , " / " , len(files) , " files.")
    
    fmt.Println("Removing record directory " + recordDir + "...")
    err = os.RemoveAll(recordDir)
    if err != nil {
      panic(err)
    }  
    fmt.Println("Removed record directory " + recordDir)
}

func main() {
    go server()
    
    scrape()
    
	readDirectory(savePath)

    ticker := time.NewTicker(60 * 60 * 24 * time.Second)
    
    for {
        select {
            case <- ticker.C: 
                scrape()
                readDirectory(savePath)
        }
    }

	fmt.Println("server end")
}
