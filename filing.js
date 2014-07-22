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

function createFiling(data) {
	var lobbyists = [];

	//keep duplicates
	/*
	data.Lobbyist.forEach(function(entry) {
		lobbyists.push(new Lobbyist(entry.FirstName, entry.LastName));
	})
	*/
	test = data;

	//get rid of duplicates and check if no lobbyists on file
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

	return new Filing(data.OrganizationName, data.ClientName, data.HouseID, data.SenateID, lobbyists);
}

function parseData(data) {
	var array = [];

	if (data.array) {
		data = data.array;

		data.forEach(function(entry) {
			array.push(createFiling(entry));
		});
	} else {
		console.log("Error parsing obj:\n" + data);
	}

	return array;
}