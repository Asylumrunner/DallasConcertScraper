package main

import (
  "strings"
  "github.com/aws/aws-lambda-go/lambda"
  "fmt"
)

func main() {
  lambda.Start(ScrapeAndParse)
}

func getVenues() []string {
  venue_list_b := DownloadFromS3("venues.txt")

  venue_list := strings.Split(venue_list_b, ",")
  act_list := make([]string, 0)
  for _, venue := range venue_list {
    act_list = append(act_list, scrape(venue)...)
  }
  return act_list


}

func ScrapeAndParse() {
  venues := getVenues()
  fmt.Println(venues)
}
