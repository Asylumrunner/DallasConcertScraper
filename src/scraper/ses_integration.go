package main

import {
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"log"
}

func SendEmail(email_body string) bool {
	svc := ses.New(session.New())
	log.Print("Established SES session")

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String("asylumrunner@gmail.com")},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(email_body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Your Scraped Concert Information Is Ready"),
			},
		},
		Source: aws.String("asylumrunner@gmail.com"),
		ReplyToAddresses: []*string{
			aws.String(asylumrunner@gmail.com),
		},
	}

	_, err := svc.SendEmail(input)

	if err != nil {
		log.Print("An error occurred trying to send an email")
    	log.Print(err.Error())
    	return false
	}

	log.Print("Email sent successfully")
	return true
}