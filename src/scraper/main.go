package main

import (
  "strings"
  "github.com/aws/aws-lambda-go/lambda"
  "log"
)

func main() {
  lambda.Start(ScrapeAndParse)
}

func getShows() []show {
  log.Print("Downloading venue list from S3!")
  venue_list_b := DownloadFromS3("venues.txt")
  log.Print("Venue list downloaded")

  log.Print("Formatting venue list and preparing for scraping")
  venue_list := strings.Split(venue_list_b, ",")
  act_list := make([]show, 0)

  var trimmed_venue string
  for _, venue := range venue_list {
    trimmed_venue = strings.TrimSpace(venue)
    log.Print("Scraping for " + trimmed_venue)
    act_list = append(act_list, scrape(trimmed_venue)...)
    log.Print("Scraping for " + trimmed_venue + " completed")
  }
  return act_list
}

func ScrapeAndParse() {
  log.Print("Lambd function has spun up!")
  shows := getShows()
  formatted_show_document := FormatScrapedData(shows)
  log.Print(formatted_show_document)
  uploaded_successfully := UploadToS3(formatted_show_document)
  if uploaded_successfully {
    log.Print("File uploaded to S3. Ceasing execution")
  }
}
