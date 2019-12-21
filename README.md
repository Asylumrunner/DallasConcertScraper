# DallasConcertScraper
A Go-based lambda function designed to scrape information about concerts in my area

## Overview
I like to see live music, and even though Dallas doesn't have the best live music scene in the world, we still get a good amount of really good shows. The problem is that there's no easy way to see what's playing, especially since there are approximately nine thousand concert/event aggregation services out there, and it seems like figuring out which shows are listed on which service is a complete dice roll. So, I decided to write myself a scraper that would periodically scrape the online calendars of concert venues I like, collect the data on who's playing, and dump all of that aggregated information to a file for me to browse through.

AWS Lambda was a fairly obvious fit for this project. Concerts don't tend to just pop out of thin air, they're usually posted months in advance, so a Lambda function set to a cron trigger could spin itself up every month or so, aggregate upcoming concerts, format them all nicely, and dump them into an S3 object for perusal at my next convenience. Hypothetically I could even email this file to myself whenever it was ready.

Go as a language choice is somewhat artificial. Python could easily handle this task just as easily, if not more so, but I've been interested in Go and looking for an excuse to use it, and with Go's adoption as an officially-supported language for AWS Lambda in 2018, it doesn't seem too big a stretch.

## Important Note
This web scraper, like most, is extremely brittle, heavily dependent on the specific HTML and CSS styling used by the websites it scraped. When first created, this scraper primarily leveraged the fact that many of Dallas's concert venues appeared to have used the same web developer, resulting in CSS stylesheets that were nearly identical.

Since this project was completed, many of these venues have since redesigned their websites (which is a good thing, most of them were pretty bad, it's one of the reasons I made this thing in the first place), and as a result have temporarily broken my ability to scrape them. 

Still, though, this project served as an interesting learning experience, and maybe as I move around I'll be able to reimplement it, so I'm leaving it around.

## Installation Steps
Prerequisites:
* An Amazon Web Services Account
* Ideally two, minimum one email address, registered to Amazon Simple Email Service
* A client ID and secret for the Spotify API

1.  From the root of this project, run the following command line instructions.
```
env GOOS-linux GOARCH=amd64 go build -i /tmp/main scraper
//This builds the actual project binary
zip -j main.zip /tmp/main
//AWS Lambda mainly takes code in the form of zipped binaries, so you need to zip it up
rm /tmp/main
//Technically optional, I like to clean up the old binary
```

2. With the binary, create an AWS Lambda function and call it whatever ("DallasConcertScraper" is fine). Mark the language as Go 1.x. When given the option, select "Upload a .zip", upload the zip file created above, and save the function.

3. In the Environment variables window, you're going to need to set environment variables with the following keys and values:
```
client_id - Your Spotify API client ID
client_secret -  Your Spotify API client secret
dest_email - The email address you want the scraped report sent to
send_email - The email address you want to send the scraped report from. This can be the same as dest_email, but most email clients will flag this as potential fraud
```

4. Create a file called venues.txt, containing every venue URL you want to scrape as a comma separated list, no spaces. Upload it to S3.

5. This scraper is written to be triggered once a month by a Cloudwatch Event, which you can set up in the Lambda console. Alternatively, if you want to trigger this manually, you can set up a Test Event in the Lambda console with whatever payload (empty JSON object is what I use), so that you can run the function just by hitting "Test"


