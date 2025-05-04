package database

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func writeToFile(fileName string, data map[string]any) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %w", err)
	}
	defer file.Close()

	for k, v := range data {
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

func readFile(fileName string) (map[string]any, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, fs.ModeTemporary)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[string]any)

	for {
		sizeBuf := make([]byte, SIZE_BYTE)
		_, err = io.ReadFull(file, sizeBuf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("error reading size: %w", err)
		}

		dataSize := binary.BigEndian.Uint32(sizeBuf)
		if dataSize == 0 {
			continue
		}

		dataBuf := make([]byte, dataSize)
		_, err = io.ReadFull(file, dataBuf)
		if err != nil {
			return nil, fmt.Errorf("error reading data: %w", err)
		}

		var kv map[string]any
		if err = json.Unmarshal(dataBuf, &kv); err != nil {
			return nil, fmt.Errorf("error unmarshalling json: %w", err)
		}

		for k, v := range kv {
			data[k] = v
		}
	}

	return data, nil
}
