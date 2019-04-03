package main

import (
  "strings"
  "github.com/aws/aws-lambda-go/lambda"
  "log"
)

func main() {
  lambda.Start(ScrapeAndParse)
}

func getVenues() []string {
  log.Print("Downloading venue list from S3!")
  venue_list_b := DownloadFromS3("venues.txt")
  log.Print("Venue list downloaded")

  log.Print("Formatting venue list and preparing for scraping")
  venue_list := strings.Split(venue_list_b, ",")
  act_list := make([]string, 0)

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
  log.Print("Lambda function has spun up!")
  venues := getVenues()
  for _, act := range venues {
    log.Print("Act found: " + act)
  }
}
