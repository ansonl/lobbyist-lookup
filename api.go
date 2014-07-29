package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Lobbyist struct {
	FirstName string `xml:"lobbyistFirstName"`
	LastName  string `xml:"lobbyistLastName"`
}

type Registration struct {
	OrganizationName string `xml:"organizationName"`
	ClientName       string `xml:"clientName"`
	SenateID         string `xml:"senateID"`
	HouseID          string `xml:"houseID"`
	//ReportYear string `xml:"reportYear"`
	//ReportType string `xml:"reportType"`
	Lobbyist []Lobbyist `xml:"alis>ali_info>lobbyists>lobbyist"` //apparently house changed their xml format on 6/10??
	//Lobbyist []Lobbyist `xml:"lobbyists>lobbyist"`
}

var rArray []Registration

var counter = 0

var startTime = time.Now()

func ExtendStringSlice(slice []string, element string) []string {
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

func ExtendResultSlice(slice []Registration, element Registration) []Registration {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]Registration, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

//return a json formatted string for a day struct
func (reg Registration) JSONString() []byte {
	b, err := json.Marshal(reg)
	if err != nil {
		panic(err)
	}
	return b
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./pages/api.html")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(data))
}

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	diff := time.Since(startTime)

	fmt.Fprintf(w, "Uptime:\t"+diff.String()+"\n\tLookups served:\t"+strconv.Itoa(counter)+" ")
	fmt.Println("Uptime requested")
}

func legislationHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./pages/legislation.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(data))
}

func autoSurnameHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	//bypass same origin policy
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//firstName := r.Form["first"]
	lastName := r.Form["surname"]

	limit := 100
	count := 0

	matches := make([]string, 0)

	//surname search
	if lastName != nil && len(lastName) > 0 && lastName[0] != "" { //check if empty param (surname=) because strings.Contains will flag empty string as match
		for _, i := range rArray {
			if count < limit {
				for _, j := range i.Lobbyist {
					if j.LastName != "" {

						for _, l := range lastName {
							if count < limit {
								if strings.Contains(strings.ToLower(j.LastName), l) {
									if len(matches) > 0 {
										for n, m := range matches {

											if strings.Replace(j.LastName, " ", "", -1) != string(m) {
												fmt.Println("comparing" + strings.Replace(j.LastName, " ", "", -1) + "|" + string(m) + string(n))

												matches = ExtendStringSlice(matches, strings.Replace(j.LastName, " ", "", -1))
												count++
											}
										}
									} else {
										matches = ExtendStringSlice(matches, strings.Replace(j.LastName, " ", "", -1))
									}
								}
							}
						}

					}
				}
			}
		}
	}

	returnString, err := json.Marshal(matches)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(returnString))

}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form)

	//bypass same origin policy
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//setup return string
	returnString := "{" + `"` + "array" + `"` + ":" + "[ "

	//firstName := r.Form["first"]
	lastName := r.Form["surname"]
	organizationName := r.Form["organization"]
	clientName := r.Form["client"]

	limit := 10
	count := 0

	counter++

	matches := []Registration(nil)

	//surname search
	if lastName != nil && len(lastName) > 0 && lastName[0] != "" { //check if empty param (surname=) because strings.Contains will flag empty string as match
		tmp := make([]Registration, 0)
		if matches != nil {
			for _, i := range matches {
				for _, j := range i.Lobbyist {
					if j.LastName != "" {
						for _, l := range lastName {
							if strings.Contains(strings.ToLower(j.LastName), l) {
								tmp = ExtendResultSlice(tmp, i)
								break
							}
						}
					}
				}
			}
			matches = tmp
		} else {
			matches = make([]Registration, 0)
			for _, i := range rArray {
				if count < limit {
					for _, j := range i.Lobbyist {
						if j.LastName != "" {
							if count < limit {
								for _, l := range lastName {
									if strings.Contains(strings.ToLower(j.LastName), l) {
										matches = ExtendResultSlice(matches, i)
										count++
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}

	//organization name search
	if organizationName != nil && len(organizationName) > 0 && organizationName[0] != "" {
		tmp := make([]Registration, 0)
		if matches != nil {
			for _, i := range matches {
				for _, l := range organizationName {
					if strings.Contains(strings.ToLower(i.OrganizationName), l) {
						tmp = ExtendResultSlice(tmp, i)
						break
					}
				}
			}
			matches = tmp
		} else {
			matches = make([]Registration, 0)
			for _, i := range rArray {
				if count < limit {
					for _, l := range organizationName {
						if strings.Contains(strings.ToLower(i.OrganizationName), l) {
							matches = ExtendResultSlice(matches, i)
							count++
							break
						}
					}
				}
			}
		}
	}

	//client name search
	if clientName != nil && len(clientName) > 0 && clientName[0] != "" {
		tmp := make([]Registration, 0)
		if matches != nil {
			for _, i := range matches {
				for _, l := range clientName {
					if strings.Contains(strings.ToLower(i.ClientName), l) {
						tmp = ExtendResultSlice(tmp, i)
						break
					}
				}
			}
			matches = tmp
		} else {
			matches = make([]Registration, 0)
			for _, i := range rArray {
				if count < limit {
					for _, l := range clientName {
						if strings.Contains(strings.ToLower(i.ClientName), l) {
							matches = ExtendResultSlice(matches, i)
							count++
							break
						}
					}
				}
			}
		}
	}

	/*
							for _, i := range rArray {
							if (count < limit) {
							if (organizationName != nil) {
							for _, k := range organizationName {
							if (strings.Contains(strings.ToLower(i.OrganizationName), k)) {
							if (lastName != nil) {
							for _, j := range i.Lobbyist {
							if (j.LastName != "") {
							for _, l := range lastName {
							if (strings.Contains(strings.ToLower(j.LastName), l)) {
							matches = ExtendResultSlice(matches, i)
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
		matches = ExtendResultSlice(matches, i)
		count++
		}
		}
		}
		}
		}
		}
		}
		}
	*/

	for _, element := range matches {
		returnString += string(element.JSONString()) + ","
	}

	returnString = returnString[:len(returnString)-1]
	returnString += "]" + "}"
	fmt.Fprintf(w, returnString)
}

func server() {
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/legislation/", legislationHandler)
	http.HandleFunc("/uptime", uptimeHandler)
	http.HandleFunc("/", handler)
	http.HandleFunc("/autosurname/", autoSurnameHandler)
	//http.ListenAndServe(":8080", nil)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("listening on port " + os.Getenv("PORT"))
}

func readDirectory(recordDir string) {
	files, err := ioutil.ReadDir("./" + recordDir + "/")
	if err != nil {
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
			if strings.Contains(filepath.Ext(f.Name()), "xml") {

				//unmarshal data and put into struct array
				err = xml.Unmarshal([]byte(data), &rArray[a])
				if err != nil {
					fmt.Println("error decoding %v: %v", f.Name(), err)
					return
				}

				a++ //increment number of files successfully parsed
			}
		}

		if a%1000 == 0 {
			fmt.Println(strconv.Itoa(a) + " files read")
		}
	}

	fmt.Println("Successfully read ", a, " / ", len(files), " files.")

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
		case <-ticker.C:
			scrape()
			readDirectory(savePath)
		}
	}

	fmt.Println("server end")
}
