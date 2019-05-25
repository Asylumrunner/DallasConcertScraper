package main

import (
  "net/http"
  "time"
  "golang.org/x/net/html"
  "log"
)

type information_type int

const (
  headliner information_type = 0 + iota
  doors_time
  show_time
  openers
  date
  none
)

func scrape(url string) []show {
  log.Print("Creating HTTP Client with a 30 second timeout")
  client := &http.Client{
    Timeout: 30 * time.Second,
  }

  log.Print("Performing GET call on " + url)
  response, err := client.Get(url)
  defer response.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  body := make([]byte, 9999)
  response.Body.Read(body)
  log.Print(string(body))

  band_names := make([]show, 0)
  tokenizer := html.NewTokenizer(response.Body)
  var data_expected information_type

  end_of_document_found := false
  var t html.Token
  var new_show show

  for {
    token := tokenizer.Next()

    switch {

      case token == html.ErrorToken:
        log.Print("Found end of HTML document. Closing")
        end_of_document_found = true

      case token == html.StartTagToken:
        t = tokenizer.Token()
        if t.Data == "h1" {
          ok, value := grabAttribute(t, "class")
          if ok && value == "headliners summary" {
            data_expected = headliner
            band_names = append(band_names, new_show)
            new_show = show{}
            new_show.venue = url
          }
          if ok && value == "headliners" {
            data_expected = headliner
          }
        } else if t.Data == "h2" {
          ok, value := grabAttribute(t, "class")
          if ok && value == "supports description"{
            data_expected = openers
          } else if ok && value == "dates" {
            data_expected = date
          }
        }

      case token == html.TextToken && data_expected == headliner:
        t = tokenizer.Token()
        //band_names = append(band_names, t.Data)
        if new_show.headliner == "" {
          new_show.headliner = t.Data
        } else {
          new_show.headliner += (", " + t.Data)
        }
        log.Print("Band name " + t.Data + " found. Appending to list")
        data_expected = none

      case token == html.TextToken && data_expected == openers:
        t = tokenizer.Token()
        new_show.openers = t.Data
        data_expected = none

      case token == html.TextToken && data_expected == date:
        t = tokenizer.Token()
        new_show.date = t.Data
        data_expected = none
    }
    if end_of_document_found {
      break
    }
  }

  return band_names
}

func grabAttribute(token html.Token, attribute string) (bool, string) {
  for _, value := range token.Attr {
    if value.Key == "class" {
      return true, value.Val
    }
  }
  return false, ""
}
