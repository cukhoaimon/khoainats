package database

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestConcurrentReadWrite(t *testing.T) {
	pwd, _ := os.Getwd()
	dbDir := filepath.Join(pwd, "tmp")
	db := newSimpleDatabase(dbDir).(*worldSimplestDatabase)
	if err := db.Init(); err != nil {
		t.Fatalf("cannot init database due to error %v", err)
	}

	var wg sync.WaitGroup
	const (
		goroutines        = 1000
		writePerGoroutine = 100
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

	// run compaction
	c := NewJobCompaction(time.Hour, dbDir).(*compactor)
	if err := c.runCompact(); err != nil {
		t.Errorf("fail to run compaction %v", err)
	}

	// Assert there is at lest 1 file named with prefix SNAPSHOT
	entries, err := os.ReadDir(dbDir)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	found := false
	for _, entry := range entries {
		if strings.Contains(entry.Name(), "SNAPSHOT") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected at least 1 file with prefix SNAPSHOT to ensure compaction succeed but encounter error %v", err)
	}
}
