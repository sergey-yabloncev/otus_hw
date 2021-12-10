package memorystorage

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu     sync.RWMutex
	events map[string]*storage.Event
	logger Logger
}

type Logger interface {
	Info(msg string)
	Error(msg string)
	Warn(msg string)
	Debug(msg string)
}

func New(logger Logger) *Storage {
	return &Storage{
		events: make(map[string]*storage.Event),
		logger: logger,
	}
}

func (s *Storage) CreateEvent(_ context.Context, e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.events[e.ID] != nil {
		return errors.New("event already exist")
	}

	s.events[e.ID] = &e
	return nil
}

func (s *Storage) UpdateEvent(_ context.Context, e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.events[e.ID] == nil {
		s.logger.Info(fmt.Sprintf("event %v not found", e.ID))
		return errors.New("event not found")
	}

	s.events[e.ID] = &e
	return nil
}

func (s *Storage) DeleteEvent(_ context.Context, e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.events[e.ID] == nil {
		s.logger.Info(fmt.Sprintf("event %v not found", e.ID))
		return errors.New("event not found")
	}

	delete(s.events, e.ID)
	return nil
}

func (s *Storage) GetEvents(_ context.Context) ([]storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	events := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		events = append(events, *event)
	}
	return events, nil
}
