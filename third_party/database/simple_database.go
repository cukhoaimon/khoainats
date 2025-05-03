package database

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	SIZE_BYTE = 4
)

type worldSimplestDatabase struct {
	data      map[string]any
	compactor JobCompaction
	mu        sync.RWMutex
	fileDir   string
}

func newSimpleDatabase(fileDir string) AbstractDatabase {
	return &worldSimplestDatabase{
		data:      make(map[string]any),
		compactor: nil,
		fileDir:   fileDir,
	}
}

func (w *worldSimplestDatabase) ReadAll() map[string]any {
	w.mu.RLock()
	data := w.data
	w.mu.RUnlock()

	return data
}

func (w *worldSimplestDatabase) Write(key string, value any) error {
	w.mu.Lock()
	w.data[key] = value
	w.mu.Unlock()

	return nil
}

func (w *worldSimplestDatabase) Read(key string) any {
	w.mu.RLock()
	data := w.data[key]
	w.mu.RUnlock()

	return data
}

func (w *worldSimplestDatabase) Init() error {
	log.Println("init worldSimplestDatabase")
	if err := w.readAllFile(w.fileDir); err != nil {
		return err
	}

	log.Println("run compaction job")
	w.compactor = NewJobCompaction(24*time.Hour, w.fileDir)
	if err := w.compactor.Run(); err != nil {
		return err
	}

	return nil
}

func (w *worldSimplestDatabase) Shutdown() error {
	log.Println("shutting down worldSimplestDatabase")
	if err := writeToFile(filepath.Join(w.fileDir, time.Now().UTC().String()), w.data); err != nil {
		return err
	}

	log.Println("stopping compaction job")
	return w.compactor.Stop()
}

func (w *worldSimplestDatabase) readAllFile(fileDir string) error {
	if fileDir == "" {
		fileDir = w.fileDir
	}
	entries, err := os.ReadDir(fileDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.Mkdir(fileDir, os.ModeDir)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if err = w.readAllFile(filepath.Join(w.fileDir, entry.Name())); err != nil {
				return err
			}
		}

		file := filepath.Join(w.fileDir, entry.Name())
		data, err := readFile(file)
		if err != nil {
			return err
		}
		w.mu.Lock()
		for k, v := range data {
			w.data[k] = v
		}
		w.mu.Unlock()
	}

	return nil
}
