package core

import (
	"log"
	"os"
)

type WatcherService interface {
	OnError(err error)
	OnCreate(name string) error
	OnRemove(name string) error
	OnWrite(name string) error
	OnRename(name string) error
	OnChmod(name string) error
}

type fileStatus struct {
	created, writing bool
}

type filesService struct {
	files map[string]fileStatus
}

func NewFilesService() WatcherService {
	return &filesService{files: map[string]fileStatus{}}
}

func (s *filesService) resetFileStatus(name string) {
	if file, ok := s.files[name]; ok {
		file.created, file.writing = false, false
		s.files[name] = file
	}
}

func (s *filesService) OnCreate(name string) error {
	if _, ok := s.files[name]; ok {
		s.resetFileStatus(name)
	} else {
		s.files[name] = fileStatus{
			created: true,
			writing: false,
		}
	}
	return nil
}

func (s *filesService) OnRemove(name string) error {
	s.resetFileStatus(name)
	return nil
}

func (s *filesService) OnRename(name string) error {
	s.resetFileStatus(name)
	return nil
}

func (s *filesService) OnChmod(name string) error {
	s.resetFileStatus(name)
	return nil
}

func (s *filesService) OnError(err error) {
	log.Println(err)
}

func (s *filesService) OnWrite(name string) error {
	if file, ok := s.files[name]; ok {
		if file.created && file.writing {
			if _, err := os.Stat(name); err == nil {
				log.Println(name, "file exist")
			} else {
				log.Println(name, "file not exist")
			}
			s.resetFileStatus(name)
		} else if file.created {
			file.writing = true
			s.files[name] = file
		} else {
			s.resetFileStatus(name)
		}
	} else {
		s.resetFileStatus(name)
	}
	return nil
}
