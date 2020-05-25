package main

import (
	"github.com/andrzejd-pl/aws-s3/aws"
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("Init the program")
	contents := "This is a new file stored in the cloud"
	r := strings.NewReader(contents)
	size := int64(len(contents))

	client := aws.NewClient(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_SECRET_KEY"),
		os.Getenv("AWS_REGION"),
		os.Getenv("AWS_S3_BUCKET"),
	)

	item, err := client.Upload(r, size, "name.txt", "directory")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(item.Name())
}
