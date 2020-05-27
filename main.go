package main

import (
	"fmt"
	"github.com/andrzejd-pl/aws-s3/aws"
	"github.com/andrzejd-pl/aws-s3/core"
	"github.com/andrzejd-pl/aws-s3/filesystem"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	catchSigint(&wg)

	client := aws.NewClient(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_SECRET_KEY"),
		os.Getenv("AWS_REGION"),
		os.Getenv("AWS_S3_BUCKET"),
	)
	service := core.NewFilesService(client)
	log.Println("Start watch directory")
	watcher := filesystem.NewFilesWatcher(service, os.Getenv("WATCHED_DIRECTORY"), &wg)

	err := watcher.Watch()
	if err != nil {
		log.Println(err)
		return
	}
	defer watcher.Close()

	wg.Wait()
}

func catchSigint(wg *sync.WaitGroup) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("To stop the program press Ctrl+C")
		_ = <-c
		log.Println("Stop the program")
		wg.Done()
	}()
}
