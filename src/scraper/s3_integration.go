package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "fmt"
)

func DownloadFromS3(filename string) string {
  svc := s3.New(session.New())
  input := &s3.GetObjectInput{
    Bucket: aws.String("ConcertScraper"),
    Key:    aws.String(filename),
  }

  result, err := svc.GetObject(input)
  if err != nil {
    fmt.Println("An error occured when downloading the object from S3")
    fmt.Println(err.Error())
    return ""
  }

  var body []byte
  body = make([]byte, 9999)
  result.Body.Read(body)
  return string(body)
}
