package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
)

type MinioWrapper struct {
	Client   *minio.Client
	S3Client *minio.Client
}

var bucketName = os.Getenv("S3_DEPLOY_BUCKET_NAME")
var deployFile = os.Getenv("S3_DEPLOY_FILE_NAME")

func (client *MinioWrapper) init() {
	client.initMinioClient()
	client.initS3MinioClient()
}

func (client *MinioWrapper) initMinioClient() {

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

	client.Client = minioClient
}

func (client *MinioWrapper) initS3MinioClient() {
	var endpoint string = os.Getenv("MINIO_S3_ENDPOINT")
	var accessKeyId string = os.Getenv("MINIO_S3_ACCESS_KEY_ID")
	var secretAccessKey string = os.Getenv("MINIO_S3_SECRET_ACCESS_KEY")
	// Initialize minio client object.
	s3Client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	client.S3Client = s3Client
}


func (client *MinioWrapper) Deploy() {
	client.EnsureBucketExists()
	client.CopyDeployFile()
}

func (client *MinioWrapper) DoesDeployBucketExist() bool {
	found, err := client.S3Client.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	return found
}

func (client *MinioWrapper) EnsureBucketExists() {
	fmt.Println("ensuring deploy bucket exists")
	var doesExist bool = client.DoesDeployBucketExist()
	if !doesExist {
		client.CreateDeployBucket()
	}
	fmt.Println("ensuring deploy bucket exists")
}

func (client *MinioWrapper) CreateDeployBucket() {
	fmt.Println("creating s3 bucket")
	// create deploy bucket
	client.S3Client.MakeBucket(context.Background(), bucketName,
		minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	fmt.Println("s3 bucket created")
}

func (client *MinioWrapper) CopyDeployFile() {
	fmt.Println("copying deployment files locally")
	object, err := client.Client.GetObject(context.Background(), bucketName, deployFile, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	localFile, err := os.Create(deployFile)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = io.Copy(localFile, object); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("deploy file copied locally")
}
