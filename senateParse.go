package main

//http://play.golang.org/p/kaZrQ2HJas

import (
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Registrant struct {
	RegistrantName    string `xml:",attr"`
	RegistrantID      string `xml:",attr"`
	RegistrantCountry string `xml:",attr"`
}
type Client struct {
	ClientName        string `xml:",attr"`
	ClientID          string `xml:",attr"`
	ContactFullname   string `xml:",attr"`
	IsStateOrLocalGov bool   `xml:",attr"`
	ClientCountry     string `xml:",attr"`
}
type Lobbyist struct {
	LobbyistName string `xml:",attr"`
}
type Filing struct {
	ID         string     `xml:"ID,attr"`
	Client     Client     `xml:Client"`
	Registrant Registrant `xml:"Registrant"`
	Lobbyists  []Lobbyist `xml:"Lobbyists>Lobbyist"`
}
type PublicFilings struct {
	Filings []Filing `xml:"Filing"`
}

func convertEncoding(input []byte) []byte {
	reader, err := charset.NewReader("utf16", strings.NewReader(string(input)))
	if err != nil {
		log.Fatal(err)
	}
	output, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func parseSenateFilings() []Filing {

	files, err := ioutil.ReadDir(savePathSenate)
	if err != nil {
		panic(err)
	}

	allSenateFilings := make([]Filing, len(files))

	fmt.Println("Reading " + strconv.Itoa(len(files)) + " files from " + savePathSenate + "...")

	a := 0 //counter for number of files successfully read

	for _, f := range files {

		oneFile := PublicFilings{}

		data, err := ioutil.ReadFile(savePathSenate + f.Name())
		if err != nil {
			fmt.Println("error reading", f.Name(), err)
			continue
		} else {
			if strings.Contains(filepath.Ext(f.Name()), "xml") {
				data = convertEncoding(data)

				data = []byte(strings.Replace(string(data), "UTF-16", "UTF-8", -1))

				if err := xml.Unmarshal([]byte(data), &oneFile); err != nil {
					fmt.Println(f.Name(), err)
				}

				for _, t := range oneFile.Filings {
					allSenateFilings = append(allSenateFilings, t)
					a++
					if a%1000 == 0 {
						fmt.Println(strconv.Itoa(a), "Senate filings read")
					}
				}
			}
		}
	}

	fmt.Println("Successfully read ", a, "Senate filings from", len(files), " files.")

	fmt.Println("Removing record directory " + savePathSenate + "...")
	err = os.RemoveAll(savePathSenate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Removed record directory " + savePathSenate)

	return allSenateFilings
}
