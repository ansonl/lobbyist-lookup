<h3>Unified Congress Lobbyist Disclosure Scrapper and Lookup</h3>

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy?template=https://github.com/ansonl/lobbyist-lookup)

- <h4>Record Retrieval</h4>
  - Latest current year House lobby disclosure filings available on [House.gov](http://disclosures.house.gov/).
    - Using the [webbrowser based search](http://disclosures.house.gov/ld/ldsearch.aspx) may result in
      > Cannot download more than 2000 records. Please refine search.
    - Using *past filings* download link utimately leads to [here](http://disclosures.house.gov/ld/LDDownload.aspx?KeepThis=true) to download filings in xml format.
      - The house.gov site uses an input element with method of POST to an asp page to serve the archive files. The site also runs on ASP which has ViewState and EventValidation enforced to prevent CSRF. ViewStateand EventValidation makes programmatic POST requests more complicated as we need to have valid ViewState and EventValidation values in order to send a valid POST request.
        - This Go program retrieves a response from the ASP server with a GET request. After parsing the hidden ViewState and EventValidation input values, we are able to construct a valid POST request which the ASP server replies back with a file stream. We write the file stream to a defined file.  
          - `houseRetrieve.go` uses `code.google.com/p/go.net/html` package to parse HTML for tokens.
          - `houseRetrieve.go` contains the archive downloading portion of the code and can be repurposed to send/received requests with other ASP sites using CSRF protection.
    - `XXXX Registration` archives contain new registrations for that year. `XXXX N Quarter` archives contain filings due for *N* quarter.
      - This program will download all archives for the current year.
  - Use predicted file naming convention for Senate filings on [Senate.gov](http://www.senate.gov/legislative/Public_Disclosure/LDA_reports.htm).
    - Senate provides xml files with up to 1000 filings per file.
      - XML files are in UTF-16 and Go expects UTF-8
        - Used `code.google.com/p/go-charset/charset` to convert UTF-8 to UTF-16.

  - Misc
    - House has ~90k filings versus Senate's ~130k filings.
    - House filings are in their individual XML file versus Senate filing being 1000 per file
    - Senate filings therefore parse faster funnily enough.

  - Retrieves lobbyist filings every day.
    - Heroku cycles dynos every 24 hrs so that also refreshes the list as well ;)

| Parameter | Comment |
| :--- | :--- |
| `__VIEWSTATE` | extracted token |
| `__EVENTVALIDATION` | extracted token |
| `selFilesXML` | requestd archive filename from page HTML input element|
| `btnDownloadXML` | needed to tell ASP to serve file? |


- <h4>Record Processing</h4>
  - XML files are then parsed into structs with `encoding/json` and held in memory for lookup.

- <h4>Running Lobbyist Lookup</h4>
  - Run in project directory
    - `go run *.go`
  - Compile, add to Go $PATH location and run
    - `go get` then `lobbyist-lookup`

- <h4>Reference:</h4>

  - http://stackoverflow.com/questions/14746750/post-request-using-python-to-asp-net-page/14747275#14747275 *About ASP ViewState*
  - https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)_Prevention_Cheat_Sheet#Viewstate_.28ASP.NET.29 *About ASP ViewState*
  - https://godoc.org/code.google.com/p/go.net/html#Attribute *html package that is under development and not included in default install of Go*
  - https://gist.github.com/hnaohiro/4572580 *Go unzipping code* *Had to be modified due to not closing file immediately after done copying it http://stackoverflow.com/questions/24197011/go-ioutil-using-too-many-file-descriptors-leak*
  - http://www.goinggo.net/2013/10/manage-dependencies-with-godep.html *Resolving dependencies with godep, must recreate godeps directory after adding new dependencies*
