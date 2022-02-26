Scraper built in go to create json file from site https://dpsu.gov.ua/en/AT-THE-BORDER-WITH-ROMANIA/ json file is created in same folder as code source.

It uses gocolly from https://github.com/gocolly/colly for scraper library.

This is go so to initialize go mod before compiling run "go mod init nameofyourmodule" for me it's "go mod init github.com/robcampbell79/romania_scraper".

To use gocolly it's go get -u github.com/gocolly/colly/... after you've initialized go mod.

I called this file main.go to compile: go build main.go
To run: go run main.go

The compile will make an erxecutable file.
