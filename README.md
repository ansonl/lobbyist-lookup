<h1>Lobbyist Disclosure Scrapper and Lookup</h1>

- Record Retrieval
  - Latest current year lobby disclosure filings available on House.gov.
  - User url at http://disclosures.house.gov/ utimately leads to *http://disclosures.house.gov/ld/LDDownload.aspx?KeepThis=true* to download filings in xml format. 
  - Using the standard interface, you cannot download more than 2000 records at a time. I actually tried calling them to ask about an alternative electronic method of retrieval, but alas that did not work.
  - The house.gov site uses an input element with method of POST to an asp page to serve the archive files. The site also runs on ASP which has ViewState and EventValidation enforced to prevent CSRF. ViewStateand EventValidation makes programmatic POST requests more complicated as we need to have valid ViewState and EventValidation values in order to send a valid POST request. 
    - This Go program retrieves a response from the ASP server with a GET request. After parsing the hidden ViewState and EventValidation input values, we are able to construct a valid POST request which the ASP server replies back with a file stream. We write the file stream to a defined file.  
  - XXXX year Registration XML archive contains new registrations for that year while the XXXX year N Quarter archives contains filings due for that quarter.
    - This program will download all archives for the current year.
  - Retrieves lobbyist filings every day. 

- Record Processing
  - XML files are then parsed into structs with `encoding/json` and held in memory for lookup.

- Run 
  - `go run api.go scrape.go`
or 
  - `go get` `lobbyist-lookup`

Sources:
About ASP ViewState
http://stackoverflow.com/questions/14746750/post-request-using-python-to-asp-net-page/14747275#14747275
https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)_Prevention_Cheat_Sheet#Viewstate_.28ASP.NET.29

https://godoc.org/code.google.com/p/go.net/html#Attribute *html package that is under development and not included in default install of Go*
https://gist.github.com/hnaohiro/4572580 *Go unzipping code* *Had to be modified due to not closing file immediately after done copying it http://stackoverflow.com/questions/24197011/go-ioutil-using-too-many-file-descriptors-leak*