# DallasConcertScraper
A Go-based lambda function designed to scrape information about concerts in my area

## Overview
I like to see live music, and even though Dallas doesn't have the best live music scene in the world, we still get a good amount of really good shows. The problem is that there's no easy way to see what's playing, especially since there are approximately nine thousand concert/event aggregation services out there, and it seems like figuring out which shows are listed on which service is a complete dice roll. So, I decided to write myself a scraper that would periodically scrape the online calendars of concert venues I like, collect the data on who's playing, and dump all of that aggregated information to a file for me to browse through.

AWS Lambda was a fairly obvious fit for this project. Concerts don't tend to just pop out of thin air, they're usually posted months in advance, so a Lambda function set to a cron trigger could spin itself up every month or so, aggregate upcoming concerts, format them all nicely, and dump them into an S3 object for perusal at my next convenience. Hypothetically I could even email this file to myself whenever it was ready.

Go as a language choice is somewhat artificial. Python could easily handle this task just as easily, if not more so, but I've been interested in Go and looking for an excuse to use it, and with Go's adoption as an officially-supported language for AWS Lambda in 2018, it doesn't seem too big a stretch.

## Minimal Viable Product Checklist
* Aggregate concert data from a set list (heh) of concert venues
* Convert that concert data into a single, human-readable document
* Store that document in an accessible location in S3

## Pretty Nice-To-Haves
* Email the document to myself once it's created
* Allow for the easy addition of new venues
* Allow for the prioritization of a list of bands I know and like
* Provide some amount of description for bands on top of what the venue website has because those descriptions are generally garbage

## Stretch Features
* Allow for completely venue-agnostic scraping, allowing the Lambda to scrape any concert website (Unlikely. Music venue websites are, bizarrely, a wasteland of web design)
* Turn the monthly document into a web page (this isn't actually hard with S3's static hosting functionality, it's just not terribly useful at the moment)
* Try and pull setlist data because I'm one of those people that likes to spoil myself on the setlist in advance
* Add a YouTube and/or Spotify link to the band's posting so I can figure out what the hell they actually sound like
