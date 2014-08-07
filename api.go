package main

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	//"net/http"
	//"strings"
	"time"
	"sync"
)

var counter = 0

var startTime = time.Now()
/*
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
	lastName := r.Form["term"]

	limit := 200
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
										duplicateFound := false
										for _, m := range matches {

											if strings.ToLower(j.LastName) == m {
												duplicateFound = true
											}
										}
										if duplicateFound == false {
											matches = ExtendStringSlice(matches, strings.ToLower(j.LastName))
											count++
										}
									} else {
										matches = ExtendStringSlice(matches, strings.ToLower(j.LastName))
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
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(returnString))

}

func autoOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	//bypass same origin policy
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//firstName := r.Form["first"]
	organizationName := r.Form["term"]

	limit := 100
	count := 0

	matches := make([]string, 0)

	//organization name search
	if organizationName != nil && len(organizationName) > 0 && organizationName[0] != "" {
		for _, i := range rArray {
			if count < limit {
				for _, l := range organizationName {
					if strings.Contains(strings.ToLower(i.OrganizationName), l) {
						if len(matches) > 0 {
							duplicateFound := false
							for _, m := range matches {

								if strings.ToLower(i.OrganizationName) == m {
									duplicateFound = true
								}
							}
							if duplicateFound == false {
								matches = ExtendStringSlice(matches, strings.ToLower(i.OrganizationName))
								count++
							}
						} else {
							matches = ExtendStringSlice(matches, strings.ToLower(i.OrganizationName))
						}
						break
					}
				}
			}
		}
	}

	returnString, err := json.Marshal(matches)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(returnString))

}

func autoClientHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	//bypass same origin policy
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//firstName := r.Form["first"]
	clientName := r.Form["term"]

	limit := 100
	count := 0

	matches := make([]string, 0)

	//organization name search
	if clientName != nil && len(clientName) > 0 && clientName[0] != "" {
		for _, i := range rArray {
			if count < limit {
				for _, l := range clientName {
					if strings.Contains(strings.ToLower(i.ClientName), l) {
						if len(matches) > 0 {
							duplicateFound := false
							for _, m := range matches {

								if strings.ToLower(i.ClientName) == m {
									duplicateFound = true
								}
							}
							if duplicateFound == false {
								matches = ExtendStringSlice(matches, strings.ToLower(i.ClientName))
								count++
							}
						} else {
							matches = ExtendStringSlice(matches, strings.ToLower(i.ClientName))
						}
						break
					}
				}
			}
		}
	}

	returnString, err := json.Marshal(matches)
	if err != nil {
		fmt.Println(err)
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

	limit := 100
	count := 0

	counter++

	matches := []Registration(nil)

	//surname search
	if lastName != nil && len(lastName) > 0 && lastName[0] != "" { //check if empty param (surname=) because strings.Contains will flag empty string as match

		if matches != nil {
			tmp := make([]Registration, 0)
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
		if matches != nil {
			tmp := make([]Registration, 0)
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
		if matches != nil {
			tmp := make([]Registration, 0)
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
	http.HandleFunc("/auto/surname/", autoSurnameHandler)
	http.HandleFunc("/auto/organization/", autoOrganizationHandler)
	http.HandleFunc("/auto/client/", autoClientHandler)
	//http.ListenAndServe(":8080", nil)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("listening on port " + os.Getenv("PORT"))
}
*/


func main() {
	//go server()

	var wg sync.WaitGroup
	wg.Add(2)

	var senateFilingArray []SenateFiling
	var houseFilingArray []HouseFiling

	go func() {
		senateFilingArray = parseSenateFilings(downloadSenateData(), &wg)
	}()
	go func() {
		houseFilingArray = parseHouseFilings(downloadHouseData(), &wg)
	}()

	wg.Wait()

	fmt.Println("Both Congress branches downloaded and parsed")

	fmt.Println(senateFilingArray, houseFilingArray)

	ticker := time.NewTicker(60 * 60 * 24 * time.Second)

	for {
		select {
		case <-ticker.C:
			//scrape()
			//readDirectory(savePathHouse)
		}
	}

	fmt.Println("server end")
}
