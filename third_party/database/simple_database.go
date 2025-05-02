package database

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	SIZE_BYTE = 4
	FILE_DIR  = "./tmp"
)

type worldSimplestDatabase struct {
	data      map[string]any
	compactor JobCompaction
	mu        sync.RWMutex
}

func newSimpleDatabase() AbstractDatabase {
	return &worldSimplestDatabase{
		data:      make(map[string]any),
		compactor: nil,
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
	if err := w.readAllFile(dbConfig{dir: FILE_DIR}); err != nil {
		return err
	}

	log.Println("run compaction job")
	w.compactor = NewJobCompaction(24*time.Hour, FILE_DIR)
	if err := w.compactor.Run(); err != nil {
		return err
	}

	return nil
}

func (w *worldSimplestDatabase) Shutdown() error {
	log.Println("shutting down worldSimplestDatabase")
	if err := writeToFile(filepath.Join(FILE_DIR, time.Now().UTC().String()), w.data); err != nil {
		return err
	}

	log.Println("stopping compaction job")
	return w.compactor.Stop()
}

type dbConfig struct {
	dir string
}

func (w *worldSimplestDatabase) readAllFile(cfg dbConfig) error {
	entries, err := os.ReadDir(cfg.dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if err = w.readAllFile(dbConfig{dir: filepath.Join(cfg.dir, entry.Name())}); err != nil {
				return err
			}
		}

		file := filepath.Join(cfg.dir, entry.Name())
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
