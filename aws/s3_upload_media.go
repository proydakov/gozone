package main

import (
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: cmd path/to/file")
	}

	m := make(map[string]string)
	m[".png"] = "image/png"
	m[".jpg"] = "image/jpeg"
	m[".jpeg"] = "image/jpeg"
	m[".mp4"] = "video/mp4"
	m[".webm"] = "video/webm"

	filepath := os.Args[1]

	ext := path.Ext(filepath)
	ctype := m[ext]
	if ctype == "" {
		log.Fatalf("Invalid file ext: '%s'", ext)
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	bucket := "tankzter-storage"
	key := path.Base(filepath)

	// The session the S3 Uploader will use
	s3Svc := s3.New(session.New(&aws.Config{Region: aws.String("eu-central-1")}))

	uploader := s3manager.NewUploaderWithClient(s3Svc)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:        file,
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: &ctype,
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}

	log.Println("Successfully uploaded to", result.Location)
}
