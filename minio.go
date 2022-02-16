package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"net/url"
	"os"
	"time"
)

var bucketName = os.Getenv("S3_DEPLOY_BUCKET_NAME")

func initMinioS3Client() *minio.Client {

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(minioClient)

	return minioClient
}

func CreatePresignedUrl(objectName string) (presignedUrlStr string, err error) {
	if objectName[0] == '/' {
		objectName = objectName[1:]
	}
	var s3Client *minio.Client = initMinioS3Client()

	duration, err := time.ParseDuration("48h")
	if err != nil {
		log.Fatalln(err)
	}
	urlValues := url.Values{}
	presignedUrl, err := s3Client.PresignedGetObject(context.Background(), bucketName, objectName, duration, urlValues)
	if err != nil {
		log.Fatalln(err)
	}
	var presignedUrlString = presignedUrl.Scheme + "://" + presignedUrl.Host + "/" + bucketName + "/" + objectName + "?" + presignedUrl.RawQuery
	return presignedUrlString, nil
}
