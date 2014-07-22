function formatLobbyist(lobbyist) {
	var tmp = '';

	lobbyist.forEach(function(entry) {
		tmp += entry.firstName + ' ' + entry.lastName + '<br>';
	});

	return tmp;
}

function formatFiling(filing, i) {
	var tmp = '';

	tmp += '<tr>';

	tmp += '<th>' + i + '</th>';
	tmp += '<th>' + filing.organization + '</th>';
	tmp += '<th>' + filing.client + '</th>';
	tmp += '<th>' + filing.houseId + '</th>';
	tmp += '<th>' + filing.senateId + '</th>';
	tmp += '<th>' + formatLobbyist(filing.lobbyistArray) + '</th>';

	tmp += '</tr>';

	return tmp;
}

function formatTableHeader() {
	return '<thead><tr><th>#</th><th>Filing Organization</th><th>Organization Client</th><th>House ID</th><th>Senate ID</th><th>Lobbyists</th></thead>';
}

function formatNoResults() {
	var tmp = '';

	tmp += '<tr>';

	tmp += '<td colspan="6" id="noFilingsRow">' + 'No filings matched' + '</td>';

	tmp += '</tr>';

	return tmp;
}

function createTable(filingArray) {
	var tmp = '';

	tmp += '<table class="table table-bordered table-hover"><colgroup><col width="auto"/><col width="200em"/><col width="200em"/><col width="auto"/><col width="auto"/><col width="auto"/></colgroup>' + formatTableHeader() + '<tbody>';

	if (filingArray.length == 0) {
		tmp += formatNoResults();
	} else {
		var i = 1;
		filingArray.forEach(function(entry) {
			tmp += formatFiling(entry, i);
			i++;
		});
	}

	tmp += '</tbody></table>';

	return tmp;
}