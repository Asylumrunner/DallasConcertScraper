package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "log"
)

func DownloadFromS3(filename string) string {
  svc := s3.New(session.New())
  log.Print("Established S3 session")
  input := &s3.GetObjectInput{
    Bucket: aws.String("concertscraper"),
    Key:    aws.String(filename),
  }

  log.Print("Getting object " + filename + " from S3")
  result, err := svc.GetObject(input)
  if err != nil {
    log.Print("An error occured when downloading the object from S3")
    log.Print(err.Error())
    return ""
  }
  log.Print("Object recieved successfully")

  var body []byte
  body = make([]byte, 9999)
  result.Body.Read(body)

  log.Print("Object read successfully: " + string(body))
  return string(body)
}
