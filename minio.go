package main

import (
	"context"
	"fmt"
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

// aws https://test-minio-bucket1.s3.us-east-2.amazonaws.com/index.html?response-content-disposition=inline&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEHkaCXVzLWVhc3QtMSJHMEUCIQCKO0Szkzl8KzRcJ7zpXsMMZVCDHZWCzo%2F737AqgF%2FjPgIgVIok1%2BCuM9OLsDCFkGZHwnuygKvWgnDAL9rNFtJe4gkq7QIIwf%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FARABGgwzNjk4NDA0OTM2MDciDKUbGScvV%2B7R3HyGPyrBAowDbUqmuJemgJC1r77%2FE2SFK8hIx2oYrzeBIGf4m%2FAhRC5FnYrMrEY28W%2FlFff72u4vPeRIASzQ1IhRndwI1EaDpZED6gJ0nGgK6jMP3vLcgi5cjvMsI5uuG3NRpciJtowyhpUs%2FXOrPjlgvCsMDrezjCsUcmUMlaRx5aLbSAB%2FQAV3DhvdihP2qxNisWZWwJ2fJ5s99KJt60YSjhBBKM9VhruTlyvdl%2BvarTAPGbkQZLpm%2FkgPYlHjE%2B9PfyPgc8VRDT4VuWYFHCdv7cczwV9LV0Narem60duT8BDK%2Bs7QmtP05HF85FEV3sbOWzxXx46QQGQxQ1d3zykNczzlcLjnjVvpvi15QicVOouUxzguh1avRyukx8koz4TWJguLXfy0TghYs6FlvirxdWhBmE2DqkTf%2Bu%2B4bGk5h%2FcPbiWBRzC%2Bx7SQBjqzAq3FsEG%2Bo3nFJqzXN4bpWfBAYahrVBWzziSBFQYWBlBk2QphyOFwiajLGjWwgg%2B9vcQMOnftcF%2BTkgILf5DfoA9Qbw1oNUDMY0J1yp36M2QFRrVkCO2ZpdDGstqn1UItlT8zsiJ4MBsBW9H7fHCsbo85X7J2AlhyODFM1VzFf67kHWQxFX40UNpIbDpF81Tx%2Fj3tDXWmDBadPkuuB00RFeoHYUU6vzzbhyyQroctUT1O9%2FSbtNPlBERfkbpbHUCyqXrsDJfwL%2BNjs%2FJdIneRBo0eaYdMRFveIXJXINIYSac4ldspFR08dhOeTRxoCBg9e2ftOb9dqLCCldpxHZ4cDkqYxY6i26mNcgFt2AWaArhmLiBYbSxB6Y81944NLrs4jisc36vMyUhDAbPkh1I0qZeM03M%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20220216T173204Z&X-Amz-SignedHeaders=host&X-Amz-Expires=240&X-Amz-Credential=ASIAVMHBWCQT7DV2GB5M%2F20220216%2Fus-east-2%2Fs3%2Faws4_request&X-Amz-Signature=582a672d28c44c5e5f51f970b8a39e5f1ca789458bcd7f40caa8a373e38e9eb0
//mine https://test-minio-bucket1.s3.us-east-2.amazonaws.com/test-minio-bucket1/index.html?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVMHBWCQTQZDEAYDY%2F20220216%2Fus-east-2%2Fs3%2Faws4_request&X-Amz-Date=20220216T173100Z&X-Amz-Expires=172800&X-Amz-SignedHeaders=host&X-Amz-Signature=b46aa9ffe65f33a9dc410ba0c83cb8c7d1ee4e05cbe81c7da268f3a09a1f65b1
func CreatePresignedUrl(objectName string) (presignedUrlStr string, err error) {
	if objectName[0] == '/' {
		objectName = objectName[1:]
	}
	fmt.Println("object trying to retrieve from s3 bucket: " + objectName)
	var s3Client *minio.Client = initMinioS3Client()

	duration, err := time.ParseDuration("48h")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("DURATION IS " + duration.String())
	}
	urlValues := url.Values{}
	log.Println("BUCKET NAME: " + bucketName)
	log.Println("OBJECT NAME: " + objectName)
	presignedUrl, err := s3Client.PresignedGetObject(context.Background(), bucketName, objectName, duration, urlValues)
	log.Println("returned presigned url is " + presignedUrl.Path + " " + presignedUrl.Host)
	fmt.Println(fmt.Sprintf("%#v", presignedUrl))
	if err != nil {
		log.Fatalln(err)
	}
	var presignedUrlString = presignedUrl.Scheme + "://" + presignedUrl.Host + "/" + bucketName + "/" + objectName + "?" + presignedUrl.RawQuery
	return presignedUrlString, nil
}
