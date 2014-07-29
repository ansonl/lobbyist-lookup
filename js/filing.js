function Filing(organization, client, houseid, senateid, lobbyistArray) {
	this.organization = organization;
	this.client = client;
	this.houseId = houseid;
	this.senateId = senateid;
	this.lobbyistArray = lobbyistArray;
}

function Lobbyist(firstName, lastName) {
	this.firstName = firstName;
	this.lastName = lastName;
}

var test;
var test2;

function createFiling(data, array) {
	var lobbyists = [];

	//keep duplicates
	/*
	data.Lobbyist.forEach(function(entry) {
		lobbyists.push(new Lobbyist(entry.FirstName, entry.LastName));
	})
	*/
	test = data;

	//get rid of duplicate lobbyists and check if no lobbyists on file
	if (data.Lobbyist !== null) {
		for (var i = 0; i < data.Lobbyist.length; i++) {
			var tmp = new Lobbyist(data.Lobbyist[i].FirstName, data.Lobbyist[i].LastName);
			var duplicateFound = false;
			for (var j = 0; j < lobbyists.length; j++) {
				if (lobbyists[j].firstName == tmp.firstName && lobbyists[j].lastName == tmp.lastName) {
					duplicateFound = true;
					break;
				}

			}
			if (duplicateFound === false) {
				lobbyists.push(tmp);
			}
		}
	} else {
		//no lobbyists
	}

	var duplicateFilingFound = false;
	array.forEach(function(entry) {

		if (entry.organization == data.OrganizationName && entry.client == data.ClientName && entry.houseId == data.HouseID && entry.senateId == data.SenateID && ($(entry.lobbyistArray).not(lobbyists).length == 0 && $(lobbyists).not(entry.lobbyistArray).length == 0 || JSON.stringify(entry.lobbyistArray) == JSON.stringify(lobbyists)) ) { //.not() may not detect some duplicates due to objects, so also OR for json string comparison
			duplicateFilingFound = true;
		}
	});

	if (duplicateFilingFound === false) {
		array.push(new Filing(data.OrganizationName, data.ClientName, data.HouseID, data.SenateID, lobbyists));
	}
}

var test;

function parseData(data) {
	var array = [];

	if (data.array) {
		data = data.array;

		data.forEach(function(entry) {
			createFiling(entry, array);
			test = data;
		});
	} else {
		console.log("Error parsing obj:\n" + data);
	}

	return array;
}
