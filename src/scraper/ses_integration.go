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
		
	}
}