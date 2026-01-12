package store

import (
	"sync"

	"github.com/aykutterzi/intellilog/internal/models"
)

type LogStore interface {
	AddLog(log models.LogEntry) error
	GetLogs() ([]models.LogEntry, error)
}

type InMemoryLogStore struct {
	logs []models.LogEntry
	mu   sync.RWMutex
}

func NewInMemoryLogStore() *InMemoryLogStore {
	return &InMemoryLogStore{
		logs: make([]models.LogEntry, 0),
	}
}

func (s *InMemoryLogStore) AddLog(log models.LogEntry) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.logs = append(s.logs, log)
	return nil
}

func (s *InMemoryLogStore) GetLogs() ([]models.LogEntry, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// Return a copy
	logsCopy := make([]models.LogEntry, len(s.logs))
	copy(logsCopy, s.logs)
	return logsCopy, nil
}
