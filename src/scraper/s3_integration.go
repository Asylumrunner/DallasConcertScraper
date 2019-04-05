package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "log"
    "bytes"
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
  body = make([]byte, *result.ContentLength)
  result.Body.Read(body)

  log.Print("Object read successfully: " + string(body))
  return string(body)
}

func UploadToS3(file_body string) bool {
  svc := s3.New(session.New())
  log.Print("Established S3 session")

  output := &s3.PutObjectInput{
    Bucket: aws.String("concertscraper"),
    Key:    aws.String("Scraped_Shows.txt"),
    Body:   bytes.NewReader([]byte(file_body)),
  }

  _, err := svc.PutObject(output)
  if err != nil {
    log.Print("An error occured when uploading the object to S3")
    log.Print(err.Error())
    return false
  }
  log.Print("Object uploaded successfully")
  return true
}
