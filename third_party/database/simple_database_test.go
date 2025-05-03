package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestConcurrentReadWrite(t *testing.T) {
	pwd, _ := os.Getwd()
	db := newSimpleDatabase(filepath.Join(pwd, "tmp")).(*worldSimplestDatabase)
	if err := db.Init(); err != nil {
		t.Fatalf("cannot init database due to error %v", err)
	}

	var wg sync.WaitGroup
	const (
		goroutines        = 100
		writePerGoroutine = 10
	)

	// Concurrent writes
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < writePerGoroutine; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				value := fmt.Sprintf("value-%d-%d", id, j)
				if err := db.Write(key, value); err != nil {
					t.Errorf("fail to write %v", err)
					return
				}
			}
		}(i)
	}

	wg.Wait()

	// Concurrent reads
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < writePerGoroutine; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				expected := fmt.Sprintf("value-%d-%d", id, j)
				found := db.Read(key)
				if found != expected {
					t.Errorf("expected %s, got %v", expected, found)
				}
			}
		}(i)
	}

	wg.Wait()

	if err := db.Shutdown(); err != nil {
		t.Fatalf("Shutdown failed: %v", err)
	}
}
