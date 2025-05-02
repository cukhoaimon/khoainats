package database

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
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
	data map[string]any
	mu   sync.RWMutex
}

func newSimpleDatabase() AbstractDatabase {
	return &worldSimplestDatabase{
		data: make(map[string]any),
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
	return w.readAllFile(dbConfig{dir: FILE_DIR})
}

func (w *worldSimplestDatabase) Shutdown() error {
	log.Println("shutting down worldSimplestDatabase")
	return w.writeToFile(dbConfig{dir: FILE_DIR})
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
		if err = w.readFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (w *worldSimplestDatabase) readFile(fileName string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.data == nil {
		w.data = make(map[string]any)
	}

	file, err := os.OpenFile(fileName, os.O_RDONLY, fs.ModeTemporary)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		sizeBuf := make([]byte, SIZE_BYTE)
		_, err = io.ReadFull(file, sizeBuf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("error reading size: %w", err)
		}

		dataSize := binary.BigEndian.Uint32(sizeBuf)
		if dataSize == 0 {
			continue
		}

		dataBuf := make([]byte, dataSize)
		_, err = io.ReadFull(file, dataBuf)
		if err != nil {
			return fmt.Errorf("error reading data: %w", err)
		}

		var kv map[string]any
		if err = json.Unmarshal(dataBuf, &kv); err != nil {
			return fmt.Errorf("error unmarshalling json: %w", err)
		}

		for k, v := range kv {
			w.data[k] = v
		}
	}

	return nil
}

func (w *worldSimplestDatabase) writeToFile(cfg dbConfig) error {
	fileName := filepath.Join(cfg.dir, time.Now().UTC().String())
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer file.Close()

	for k, v := range w.data {
		kv := map[string]any{k: v}
		jsonData, err := json.Marshal(kv)
		if err != nil {
			return fmt.Errorf("failed to marshal data: %w", err)
		}

		sizeBuf := make([]byte, SIZE_BYTE)
		binary.BigEndian.PutUint32(sizeBuf, uint32(len(jsonData)))
		if _, err := file.Write(sizeBuf); err != nil {
			return fmt.Errorf("failed to write size: %w", err)
		}

		if _, err := file.Write(jsonData); err != nil {
			return fmt.Errorf("failed to write json data: %w", err)
		}
	}

	return nil
}
