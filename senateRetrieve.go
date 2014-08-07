package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var savePathSenate = "./senateFilings/"
var tmpNameSenate = "tmp"
var fileExtSenate = ".zip"
var linkSenate = "http://soprweb.senate.gov/downloads/"

func generateFilenames() []string {
	//Create 2014_1 - 2014_4 and 2013_1 - 2013_4
	filenameArray := make([]string, 0)

	for i := 1; i < 5; i++ {
		filenameArray = append(filenameArray, strconv.Itoa(time.Now().Year()-1)+"_"+strconv.Itoa(i)+fileExtSenate)
		filenameArray = append(filenameArray, strconv.Itoa(time.Now().Year())+"_"+strconv.Itoa(i)+fileExtSenate)
	}

	return filenameArray
}

func downloadSenateFiles(filename string, wg *sync.WaitGroup) {
	fmt.Println("Downloading " + filename + " to " + savePathSenate + filename + "...")

	//Send GET request to download
	pres, err := http.Get(linkSenate + filename)

	robots, err := ioutil.ReadAll(pres.Body)
	pres.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(savePathSenate+filename, robots, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded " + filename + " to " + savePathSenate + filename)

	//Unzipping
	fmt.Println("Unzipping " + savePathSenate + filename + " to " + savePathSenate + "...")

	err = Unzip(savePathSenate+filename, savePathSenate)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Unzipped " + savePathSenate + filename + " to " + savePathSenate)
	}
	//Waitgroup done
	wg.Done()
}

func spawnDownloads() string {

	filenameArray := generateFilenames()

	//Check if save directory exists and create it if nonexistent
	if _, err := os.Stat(savePathSenate); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(savePathSenate, 0777)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	fmt.Println(savePathSenate + " records directory exists or was made")

	//Download each file in a thread
	var wg sync.WaitGroup
	for _, filename := range filenameArray {
		wg.Add(1)
		go downloadSenateFiles(filename, &wg)
	}
	wg.Wait()
	fmt.Println("All Senate files downloaded and unzipped.")

	return savePathSenate
}

func downloadSenateData() string {
	return spawnDownloads()

}
