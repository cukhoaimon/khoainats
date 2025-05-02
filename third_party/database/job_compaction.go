package database

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// JobCompaction represents a background task that compacts database storage.
//
// Compaction typically merges files, removes deleted or outdated entries,
// and reclaims space to improve performance and reduce storage usage.
type JobCompaction interface {
	Run() error
	Stop() error
}

func NewJobCompaction(interval time.Duration, dbDir string) JobCompaction {
	return &compactor{
		stopChan: make(chan any),
		stopped:  make(chan any),
		interval: interval,
		dbDir:    dbDir,
	}
}

type compactor struct {
	stopChan chan any
	stopped  chan any
	interval time.Duration
	dbDir    string
}

func (c *compactor) Run() error {
	go func() {
		ticker := time.NewTicker(c.interval)
		defer func() {
			close(c.stopped)
			ticker.Stop()
		}()

		for {
			select {
			case <-ticker.C:
				if err := c.runCompact(); err != nil {
					log.Printf("compaction error: %v", err)
				}
			case <-c.stopChan:
				return
			}
		}

	}()

	return nil
}

func (c *compactor) Stop() error {
	close(c.stopChan)
	<-c.stopped
	return nil
}

func (c *compactor) runCompact() error {
	entries, err := os.ReadDir(c.dbDir)
	if err != nil {
		return err
	}

	joinedData := make(map[string]any)
	for _, entry := range entries {
		file := filepath.Join(c.dbDir, entry.Name())
		data, err := readFile(file)
		if err != nil {
			return err
		}
		for k, v := range data {
			joinedData[k] = v
		}
	}

	newFile := filepath.Join(c.dbDir, time.Now().UTC().String())
	if err = writeToFile(newFile, joinedData); err != nil {
		return err
	}

	for _, entry := range entries {
		if err = os.Remove(filepath.Join(c.dbDir, entry.Name())); err != nil {
			return err
		}
	}

	return nil
}
