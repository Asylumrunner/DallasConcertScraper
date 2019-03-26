package main

import (
  "net/http"
  "time"
  "golang.org/x/net/html"
  "log"
)

func scrape(url string) []string {
  client := &http.Client{
    Timeout: 30 * time.Second,
  }

  response, err := client.Get(url)
  defer response.Body.Close()
  if err != nil {
    log.Fatal(err)
  }

  band_names := make([]string, 0)
  tokenizer := html.NewTokenizer(response.Body)
  band_name_expected := false
  var t html.Token
  for {
    token := tokenizer.Next()
    switch {
    case token == html.ErrorToken:
      break
    case token == html.StartTagToken:
      t = tokenizer.Token()
      isH1 := t.Data == "h1"
      if isH1 {
        ok, value := grabAttribute(t, "class")
        if ok && value == "headliners summary" {
          band_name_expected = true
        }
      }
    case token == html.TextToken && band_name_expected:
        t = tokenizer.Token()
        band_names = append(band_names, t.Data)
        band_name_expected = false
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
