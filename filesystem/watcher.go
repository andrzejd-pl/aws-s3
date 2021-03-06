package filesystem

import (
	"github.com/andrzejd-pl/aws-s3/core"
	"github.com/fsnotify/fsnotify"
	"log"
	"sync"
)

type Watcher interface {
	Watch() error
	Close() error
}

type filesWatcher struct {
	service   core.WatcherService
	watcher   *fsnotify.Watcher
	directory string
	waitGroup *sync.WaitGroup
}

func NewFilesWatcher(service core.WatcherService, directory string, wg *sync.WaitGroup) Watcher {
	return &filesWatcher{service: service, directory: directory, waitGroup: wg}
}

func (fw *filesWatcher) Watch() error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	fw.watcher = w

	go func() {
		for {
			select {
			case event, ok := <-fw.watcher.Events:
				if !ok {
					fw.waitGroup.Done()
					return
				}
				log.Println(event)

				var err error
				if event.Op&fsnotify.Write == fsnotify.Write {
					err = fw.service.OnWrite(event.Name)
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					err = fw.service.OnCreate(event.Name)
				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					err = fw.service.OnChmod(event.Name)
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					err = fw.service.OnRename(event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					err = fw.service.OnRemove(event.Name)
				} else if err != nil {
					fw.waitGroup.Done()
					return
				}
			case err, ok := <-fw.watcher.Errors:
				if !ok {
					return
				}
				fw.service.OnError(err)
			}
		}
	}()

	return fw.watcher.Add(fw.directory)
}

func (fw *filesWatcher) Close() error {
	if fw.watcher != nil {
		return fw.watcher.Close()
	}

	return nil
}
