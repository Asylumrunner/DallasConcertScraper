package main

func RemoveInvalidValues(shows []show) (cleaned_shows []show) {
	for _, s := range shows {
		if len(s.headliner) < 200 {
			cleaned_shows = append(cleaned_shows, s)
		}
	}
	return
}
func FormatScrapedData(shows []show) string {
	var output_string string = ""
	for _, single_show := range shows {
		output_string += ("SHOW: " + single_show.headliner + "\n")
		output_string += (single_show.date + "\n")
		if single_show.openers != "" {
			output_string += ("Openers: " + single_show.openers + "\n")
		}
		if single_show.spotify_url != "" {
			output_string += ("Spotify Link: " + single_show.spotify_url + "\n")
		}
		output_string += "\n"
	}
	return output_string
}

func FormatIntoHTMLBody(shows []show) string {
	html_email_body := ""
	html_email_body += "<h1>Band Details For Dallas Shows</h1>"
	html_email_body += "<h2>Execution Details and Full Data Output Can Be Found In AWS</h2>"

	for _, single_show := range shows {
		html_email_body += ("<h3>Show: " + single_show.headliner + "</h3>")
		html_email_body += ("<p>Date: " + single_show.date + "</p>")
		html_email_body += ("<p>Scraped From " + single_show.venue + "</p>")
		if single_show.openers != "" {
			html_email_body += ("<p>Openers: " + single_show.openers + "</p>")
		}
		if single_show.spotify_url != "" {
			html_email_body += ("<a href=\"" + single_show.spotify_url + "\">Spotify Link</a>")
		}
	}

	return html_email_body
}