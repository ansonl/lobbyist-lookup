var queryForm;
var lookupButton;

var leftDiv;
var rightDiv;

var mainContentDiv;
var formDiv;
var loadingDiv;
var resultDiv;

var lookupSuccess = function (data) {
	loadingDiv.css('display', 'none');
	resultDiv.css('display', 'inline');

	leftDiv.attr('class', 'col-md-1');
	rightDiv.attr('class', 'col-md-1');
	mainContentDiv.attr('class', 'col-md-10');

	//array of filing results returned by parseData() located in filing.js

	var filings = parseData(JSON.parse(data));

	resultDiv.html(createTable(filings));
};

var lookupError = function (data) {
	loadingDiv.css('display', 'none');
	resultDiv.css('display', 'inline');
};

var lookupComplete = function (data) {
	loadingDiv.css('display', 'none');
};

var formChanged = function (data) {
	if (data.target.value.length > 0)
		lookupButton.prop('disabled', false);
}

$(document).ready(function() {
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

	$('.form-control').on('input', formChanged);
	$('.form-control').change(formChanged);

	queryForm.submit(function() {

		if ($('#surnameInput').val().length > 0 || $('#organizationInput').val().length > 0 || $('#clientInput').val().length > 0) {

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
			console.log('not pass')
		}
		return false;
	});
})


