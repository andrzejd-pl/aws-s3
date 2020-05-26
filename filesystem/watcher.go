package filesystem

import "github.com/fsnotify/fsnotify"

type onEventFunction func(event fsnotify.Event) error
type onErrorFunction func(err error)

type Watcher interface {
	Watch() error
	Close() error
}

type filesWatcher struct {
	onError   onErrorFunction
	onEvent   onEventFunction
	watcher   *fsnotify.Watcher
	directory string
}

func NewFilesWatcher(onError onErrorFunction, onEvent onEventFunction, directory string) Watcher {
	return &filesWatcher{onError: onError, onEvent: onEvent, directory: directory}
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
					return
				}

				err := fw.onEvent(event)
				if err != nil {
					return
				}
			case err, ok := <-fw.watcher.Errors:
				if !ok {
					return
				}
				fw.onError(err)
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
