package main

import (
	"fmt"
	"strings"
	"time"
)

type GenericLobbyist struct {
	FirstName string
	LastName  string
}

type GenericFiling struct {
	OrganizationName string
	ClientName       string
	SenateID         string
	HouseID          string
	ReportYear       string
	ReportType       string
	Lobbyist         []GenericLobbyist
}

func combine(houseFilingArray []HouseFiling, senateFilingArray []SenateFiling) []GenericFiling {
	beginParseTime := time.Now()

	fmt.Println("Combining", len(houseFilingArray)+len(senateFilingArray), "filings...")

	allHouseFilings := make([]GenericFiling, len(houseFilingArray)+len(senateFilingArray))

	combinedFilingCounter := 0

	for _, houseFiling := range houseFilingArray {
		tmpFiling := GenericFiling{}
		tmpFiling.OrganizationName = houseFiling.OrganizationName
		tmpFiling.ClientName = houseFiling.ClientName
		tmpFiling.SenateID = houseFiling.SenateID
		tmpFiling.HouseID = houseFiling.HouseID
		tmpFiling.ReportYear = houseFiling.ReportYear
		tmpFiling.ReportType = houseFiling.ReportType

		tmpLobbyistArray := make([]GenericLobbyist, len(houseFiling.Lobbyist))
		for _, houseLobbyist := range houseFiling.Lobbyist {
			tmpLobbyistArray = append(tmpLobbyistArray, GenericLobbyist{houseLobbyist.FirstName, houseLobbyist.LastName})
		}
		tmpFiling.Lobbyist = tmpLobbyistArray

		combinedFilingCounter++
		if combinedFilingCounter%10000 == 0 {
			fmt.Println("Combined", combinedFilingCounter, "filings")
		}

		allHouseFilings = append(allHouseFilings, tmpFiling)
	}
	for _, senateFiling := range senateFilingArray {
		tmpFiling := GenericFiling{}
		tmpFiling.OrganizationName = senateFiling.Registrant.RegistrantName
		tmpFiling.ClientName = senateFiling.Client.ClientName
		tmpFiling.SenateID = senateFiling.ID
		//tmpFiling.HouseID = nil //no house ID provided in senate filings; because senate register before house?
		tmpFiling.ReportYear = senateFiling.Year
		tmpFiling.ReportType = senateFiling.Type

		tmpLobbyistArray := make([]GenericLobbyist, len(senateFiling.Lobbyists))
		for _, senateLobbyist := range senateFiling.Lobbyists {
			//attempt to get first and last name
			var firstName string
			var lastName string
			if strings.Index(senateLobbyist.LobbyistName, ",") < 0 {
				firstName = senateLobbyist.LobbyistName
				lastName = senateLobbyist.LobbyistName
			} else {
				firstName = senateLobbyist.LobbyistName[strings.Index(senateLobbyist.LobbyistName, ",")+1:]
				lastName = senateLobbyist.LobbyistName[:strings.Index(senateLobbyist.LobbyistName, ",")]
			}

			tmpLobbyistArray = append(tmpLobbyistArray, GenericLobbyist{firstName, lastName})
		}
		tmpFiling.Lobbyist = tmpLobbyistArray

		combinedFilingCounter++
		if combinedFilingCounter%10000 == 0 {
			fmt.Println("Combined", combinedFilingCounter, "filings")
		}

		allHouseFilings = append(allHouseFilings, tmpFiling)
	}

	fmt.Println("Done combining", len(houseFilingArray)+len(senateFilingArray), "filings in", time.Since(beginParseTime).String())

	return allHouseFilings
}
