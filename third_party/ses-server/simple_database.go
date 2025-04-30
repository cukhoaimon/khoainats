package ses_server

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

const (
	SIZE_BYTE = 4
	FILE_DIR  = "./dblog"
)

type worldSimplestDatabase struct {
	data map[string]any
}

func newSimpleDatabase() PersistenceStorage {
	return &worldSimplestDatabase{
		data: make(map[string]any),
	}
}

func (w *worldSimplestDatabase) ReadAll() map[string]any {
	return w.data
}

func (w *worldSimplestDatabase) Write(key string, value any) error {
	w.data[key] = value
	return nil
}

func (w *worldSimplestDatabase) Read(key string) any {
	return w.data[key]
}

func (w *worldSimplestDatabase) Init() error {
	log.Println("init worldSimplestDatabase")
	return w.readFromFile(dbConfig{dir: FILE_DIR})
}

func (w *worldSimplestDatabase) Shutdown() error {
	log.Println("shutting down worldSimplestDatabase")

	return w.writeToFile(dbConfig{dir: FILE_DIR})
}

type dbConfig struct {
	dir string
}

func (w *worldSimplestDatabase) readFromFile(cfg dbConfig) error {
	w.data = make(map[string]any)

	file, err := os.OpenFile(cfg.dir, os.O_RDONLY, fs.ModeTemporary)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		sizeBuf := make([]byte, SIZE_BYTE)
		_, err := io.ReadFull(file, sizeBuf)
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
		if err := json.Unmarshal(dataBuf, &kv); err != nil {
			return fmt.Errorf("error unmarshalling json: %w", err)
		}

		for k, v := range kv {
			w.data[k] = v
		}
	}

	return nil
}

func (w *worldSimplestDatabase) writeToFile(cfg dbConfig) error {
	file, err := os.OpenFile(cfg.dir, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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

		sizeBuf := make([]byte, 4)
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
