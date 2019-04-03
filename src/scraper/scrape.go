package main

import (
  "net/http"
  "time"
  "golang.org/x/net/html"
  "log"
)

func scrape(url string) []string {
  log.Print("Creating HTTP Client with a 30 second timeout")
  client := &http.Client{
    Timeout: 30 * time.Second,
  }

  log.Print("Performing GET call on " + url)
  response, err := client.Get("http://www.treesdallas.com/listing/")
  defer response.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  body := make([]byte, 9999)
  response.Body.Read(body)
  log.Print(string(body))

  band_names := make([]string, 0)
  tokenizer := html.NewTokenizer(response.Body)
  band_name_expected := false
  end_of_document_found := false
  var t html.Token
  for {
    token := tokenizer.Next()
    switch {
    case token == html.ErrorToken:
      log.Print("Found end of HTML document. Closing")
      end_of_document_found = true
    case token == html.StartTagToken:
      t = tokenizer.Token()
      isH1 := t.Data == "h1"
      if isH1 {
        log.Print("Found H1 header tag. Checking class")
        ok, value := grabAttribute(t, "class")
        if ok && value == "headliners summary" {
          log.Print("Headliner tag found. Preparing to collect band name")
          band_name_expected = true
        }
      }
    case token == html.TextToken && band_name_expected:
        t = tokenizer.Token()
        band_names = append(band_names, t.Data)
        log.Print("Band name " + t.Data + " found. Appending to list")
        band_name_expected = false
    }
    if end_of_document_found {
      break
    }
  }

  return band_names
}

func grabAttribute(token html.Token, attribute string) (bool, string) {
  log.Print("Checking HTML token for attribute " + attribute)
  for _, value := range token.Attr {
    if value.Key == "class" {
      log.Print(attribute + " attribute found!")
      return true, value.Val
    }
  }
  log.Print(attribute + " attribute not found. Returning false.")
  return false, ""
}
