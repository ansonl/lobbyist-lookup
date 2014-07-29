var queryForm;
var lookupButton;

var leftDiv;
var rightDiv;

var mainContentDiv;
var formDiv;
var loadingDiv;
var resultDiv;

var infoDiv;

var loadingText;

var newLookupButton;

var setupForNewLookup = function() {
    resultDiv.fadeOut(100, function() {
        queryForm.trigger('reset');
        mainContentDiv.attr('class', 'col-md-6');
        formDiv.fadeIn(100);
        infoDiv.fadeIn(100);
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

    infoDiv.fadeOut(100);
    loadingDiv.fadeOut(100);
    resultDiv.fadeIn(100);

    leftDiv.attr('class', 'col-md-0');
    rightDiv.attr('class', 'col-md-0');
    mainContentDiv.attr('class', 'col-md-12');

    newLookupButton = $('#newLookupButton');
    newLookupButton.click(setupForNewLookup);
};

var lookupError = function (data) {
    loadingText.fadeOut(100, function() {
        if (resultDiv.attr('display') == 'none') {
            loadingText.text('Error encountered');
            loadingText.fadeIn(100);
        }
    });

    loadingDiv.css('display', 'none');
    resultDiv.css('display', 'inline');

    var tmp = '';

    tmp += 'Lookup unsuccessful. System provided error: <br><blockquote cite="https://github.com/ansonl/lobbyist-lookup">' + data.statusText + '</blockquote>';
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
        $('#forkMeDiv').delay(1500).fadeIn(400).fadeOut(400).fadeIn(400).fadeOut(400).fadeIn(400).fadeOut(7000);

        queryForm = $('#queryForm');
        lookupButton = $('#lookupButton');

        lookupButton.prop('disabled', true);

        leftDiv = $('#leftDiv');
        rightDiv = $('#rightDiv');

        mainContentDiv = $('#mainContentDiv');
        formDiv = $('#formDiv');
        loadingDiv = $('#loadingDiv');
        resultDiv = $('#resultDiv');

        infoDiv = $('#infoDiv');

        loadingText = $('#loadingText');

        $('.form-control').on('input', formEditted);
        $('.form-control').change(formChanged);

        $( "#surnameInput" ).autocomplete({
            source: "http://lobbyist.herokuapp.com/autosurname/",
            minLength: 1,
            select: function( event, ui ) {

            }
        });

        queryForm.submit(function() {

            document.activeElement.blur();
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
