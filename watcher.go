package main

import (
	"io/fs"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func StartWatching(dir string, exceptions []string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		for _, s := range exceptions {
			if filepath.Ext(info.Name()) == s {
				return nil
			}
		}
		watcher.Add(path)
		return nil
	})

	return watcher, err
}
