package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	bucket := "tankzter-storage"
	key := "HelloS3.txt"

	svc := s3.New(session.New(&aws.Config{Region: aws.String("eu-central-1")}))
	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Println("Failed to list buckets", err)
		return
	}

	if len(result.Buckets) == 0 {
		_, err := svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: &bucket,
		})
		if err != nil {
			log.Println("Failed to create bucket", err)
			return
		}
		log.Printf("Successfully created bucket %s \n", bucket)
	}
	if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: &bucket}); err != nil {
		log.Printf("Failed to wait for bucket to exist %s, %s\n", bucket, err)
		return
	}

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: &bucket,
		Body:   strings.NewReader("Hello World!"),
		Key:    &key,
	})
	if err != nil {
		log.Printf("Failed to upload data to %s/%s, %s\n", bucket, key, err)
		return
	}

	log.Printf("Successfully upload data with key %s\n", key)
}
