function formatLobbyist(lobbyist) {
	var tmp = '';

	if (lobbyist.length === 0) {
		tmp += '<span class="noFilingsRow">' + 'No lobbyists on file' + '</span>';
	} else {
		lobbyist.forEach(function(entry) {
			if ($('#surnameInput').val() !== '' && entry.lastName.toLowerCase().indexOf($('#surnameInput').val().toLowerCase()) > -1) {
				tmp += '<span class="highlightedMatch">' + entry.firstName.toUpperCase() + ' ' + entry.lastName.toUpperCase() + '</span><br>';
			} else {
				tmp += entry.firstName.toUpperCase() + ' ' + entry.lastName.toUpperCase() + '<br>';
			}
		});
	}

	return tmp;
}

function formatGenericField(field) {
	var tmp = '';
	if (field === '') {
		tmp += '<span class="noFilingsRow">' + 'N/A' + '</span>';
	} else {
		tmp += field.toUpperCase();
	}

	return tmp;
}

function formatFiling(filing, i) {
	var tmp = '';

	tmp += '<tr>';

	tmp += '<td>' + i + '</td>';
	tmp += '<td>' + formatGenericField(filing.organization) + '</td>';
	tmp += '<td>' + formatGenericField(filing.client) + '</td>';
	tmp += '<td>' + formatGenericField(filing.houseId) + '</td>';
	tmp += '<td>' + formatGenericField(filing.senateId) + '</td>';
	tmp += '<td class="lobbyistColumn">' + formatLobbyist(filing.lobbyistArray) + '</td>';

	tmp += '</tr>';

	return tmp;
}

function formatTableHeader() {
	return '<thead><tr><th>#</th><th>Filing Organization</th><th>Organization Client</th><th>House ID</th><th>Senate ID</th><th>Lobbyists</th></thead>';
}

function formatNoResults() {
	var tmp = '';

	tmp += '<tr>';

	tmp += '<td colspan="6" class="noFilingsRow">' + 'No filings matched' + '</td>';

	tmp += '</tr>';

	return tmp;
}

function createTable(filingArray) {
	var tmp = '';

	tmp += '<h5>';
	if (filingArray.length > 9) {
		tmp += 'First ' + filingArray.length + ' filings shown. There may be more matching filings.';
	} else {
		tmp += filingArray.length + ' filings found.';
	}
	tmp += '</h5>';
	
	tmp += '<table class="table table-bordered table-hover table-condensed table-responsive"><colgroup><col width="auto"/><col width="200em"/><col width="200em"/><col width="auto"/><col width="auto"/><col width="auto"/></colgroup>' + formatTableHeader() + '<tbody>';

	if (filingArray.length === 0) {
		tmp += formatNoResults();
	} else {
		var i = 1;
		filingArray.forEach(function(entry) {
			tmp += formatFiling(entry, i);
			i++;
		});
	}

	tmp += '</tbody></table>';
	
	tmp += '<button type="button" class="btn btn-default" id="newLookupButton">New Lookup</button>';

	return tmp;
}