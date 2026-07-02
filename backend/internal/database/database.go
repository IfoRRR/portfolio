package database

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"portfolio/internal/models"
)

type store struct {
	Clicks    []clickRecord    `json:"clicks"`
	Downloads []downloadRecord `json:"downloads"`
}

type clickRecord struct {
	Link      string    `json:"link"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type downloadRecord struct {
	CreatedAt time.Time `json:"created_at"`
}

type DB struct {
	path  string
	mu    sync.Mutex
	store store
}

func Open(path string) (*DB, error) {
	db := &DB{path: path}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return db, db.save()
		}
		return nil, err
	}

	if err := json.Unmarshal(data, &db.store); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Close() error {
	return nil
}

func (db *DB) save() error {
	data, err := json.MarshalIndent(db.store, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(db.path, data, 0644)
}

func (db *DB) RecordClick(link, eventType string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.store.Clicks = append(db.store.Clicks, clickRecord{
		Link:      link,
		Type:      eventType,
		CreatedAt: time.Now().UTC(),
	})

	return db.save()
}

func (db *DB) RecordResumeDownload() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.store.Downloads = append(db.store.Downloads, downloadRecord{
		CreatedAt: time.Now().UTC(),
	})

	return db.save()
}

func (db *DB) GetStats() (*models.AnalyticsStats, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	stats := &models.AnalyticsStats{
		TotalClicks:     len(db.store.Clicks),
		ResumeDownloads: len(db.store.Downloads),
	}

	counts := map[string]models.LinkStat{}
	for _, c := range db.store.Clicks {
		key := c.Link + "|" + c.Type
		item := counts[key]
		item.Link = c.Link
		item.Type = c.Type
		item.Count++
		counts[key] = item
	}

	for _, item := range counts {
		stats.Links = append(stats.Links, item)
	}

	return stats, nil
}
