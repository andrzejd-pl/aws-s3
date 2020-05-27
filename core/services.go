package core

import (
	"github.com/andrzejd-pl/aws-s3/aws"
	"log"
	"os"
	"time"
)

type WatcherService interface {
	OnError(err error)
	OnCreate(name string) error
	OnRemove(name string) error
	OnWrite(name string) error
	OnRename(name string) error
	OnChmod(name string) error
}

type filesService struct {
	client aws.ClientInterface
}

func NewFilesService(client aws.ClientInterface) WatcherService {
	return &filesService{client: client}
}

func (s *filesService) OnCreate(string) error {
	return nil
}

func (s *filesService) OnRemove(string) error {
	return nil
}

func (s *filesService) OnRename(string) error {
	return nil
}

func (s *filesService) OnChmod(string) error {
	return nil
}

func (s *filesService) OnError(err error) {
	log.Println(err)
}

func (s *filesService) OnWrite(name string) error {
	time.Sleep(time.Second)
	file, err := os.Open(name)
	if err != nil {
		log.Printf("file name: \"%s\", error: \"%v\"\n", name, err)
		return nil
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		log.Printf("file name: \"%s\", error: \"%v\"\n", name, err)
		return nil
	}

	if fileStat.IsDir() {
		log.Printf("%s is directory", name)
		return nil
	}

	item, err := s.client.Upload(file, fileStat.Size(), file.Name(), "windows")
	if err != nil {
		log.Printf("file name: \"%s\", error: \"%v\"\n", name, err)
		return nil
	}

	log.Printf("file %s uploaded, name in aws: %s", name, item.Name())
	return nil
}
