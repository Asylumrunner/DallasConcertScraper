package main

func FormatScrapedData(shows []show) string {
	var output_string string = ""
	for _, single_show := range shows {
		output_string += ("SHOW: " + single_show.headliner + "\n")
		output_string += (single_show.date + "\n")
		if single_show.openers != "" {
			output_string += ("Openers: " + single_show.openers + "\n")
		}
		output_string += "\n"
	}
	return output_string
}