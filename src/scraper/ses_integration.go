package main

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"log"
	"os"
)

func SendEmail(email_body string) bool {
	svc := ses.New(session.New())
	log.Print("Established SES session")

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(os.Getenv("dest_email"))},
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
		Source: aws.String(os.Getenv("send_email")),
		ReplyToAddresses: []*string{
			aws.String(os.Getenv("send_email")),
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