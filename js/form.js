var queryForm;
var lookupButton;

var leftDiv;
var rightDiv;

var mainContentDiv;
var formDiv;
var loadingDiv;
var resultDiv;

var loadingText;

var newLookupButton;

var error;

var setupForNewLookup = function() {
    resultDiv.fadeOut(100, function() {
        queryForm.trigger('reset');
		formDiv.fadeIn(100);
	});
}

var lookupSuccess = function (data) {
	loadingText.fadeOut(100, function() {
        if (resultDiv.attr('display') == 'none') {
            loadingText.text('Parsing results locally');
            loadingText.fadeIn(100);
        }
	});

	//array of filing results returned by parseData() located in filing.js
	var filings = parseData(JSON.parse(data));
	resultDiv.html(createTable(filings));

	loadingDiv.css('display', 'none');
	resultDiv.css('display', 'inline');

	leftDiv.attr('class', 'col-lg-1');
	rightDiv.attr('class', 'col-lg-1');
	mainContentDiv.attr('class', 'col-lg-10');
	
	newLookupButton = $('#newLookupButton');
	newLookupButton.click(setupForNewLookup);
};

var lookupError = function (data) {
    console.log(data);
    error=data;
    
    loadingText.fadeOut(100, function() {
        if (resultDiv.attr('display') == 'none') {
		    loadingText.text('Error encountered');
		    loadingText.fadeIn(100);
        }
	});
    
	loadingDiv.css('display', 'none');
	resultDiv.css('display', 'inline');
	
	var tmp = '';
	
	tmp += 'Lookup unsuccessful. System says: <br><blockquote cite="https://github.com/ansonl/lobbyist-lookup">' + data.responseText + '</blockquote>';
	tmp += '<br><button type="button" class="btn btn-default" id="newLookupButton">New Lookup</button>';
	
	resultDiv.html(tmp);
	
	newLookupButton = $('#newLookupButton');
	newLookupButton.click(setupForNewLookup);
};

var lookupComplete = function (data) {
	loadingDiv.css('display', 'none');
};

var formEditted = function (data) {
	if (data.target.value.length > 0)
		lookupButton.prop('disabled', false);
}

var formChanged = function (data) {
	//lowercase input for matching
	data.target.value = data.target.value.toLowerCase();
}

$(document).ready(function() {
    $('#noJavascriptDiv').css('display', 'none');
    
	$('body').css('display', 'none');
	$('body').fadeIn(100);

	queryForm = $('#queryForm');
	lookupButton = $('#lookupButton');

	lookupButton.prop('disabled', true);

	leftDiv = $('#leftDiv');
	rightDiv = $('#rightDiv');

	mainContentDiv = $('#mainContentDiv');
	formDiv = $('#formDiv');
	loadingDiv = $('#loadingDiv');
	resultDiv = $('#resultDiv');

	loadingText = $('#loadingText');

	$('.form-control').on('input', formEditted);
	$('.form-control').change(formChanged);

	queryForm.submit(function() {

		if ($('#surnameInput').val().length > 0 || $('#organizationInput').val().length > 0 || $('#clientInput').val().length > 0) {

			loadingText.text('Looking up filings');
			formDiv.fadeOut(100, function() {
				//Check if ajax callback has already been called
				if (resultDiv.css('display') == 'none')
					loadingDiv.fadeIn(100);
			});
			

			var url = 'http://lobbyist.herokuapp.com/api/';
			$.ajax({
				type: 'GET',
				url: url,
				data: queryForm.serialize(),
				success: lookupSuccess,
				error: lookupError,
				complete: lookupComplete
			});
		} else {
			console.log('Fields blank')
		}
		return false;
	});
})


