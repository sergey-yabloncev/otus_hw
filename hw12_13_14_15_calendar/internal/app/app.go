package app

import (
	"context"
	"fmt"

	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type App struct { // TODO
	logger  logger.Logger
	storage Storage
}

type Storage interface {
	CreateEvent(context.Context, storage.Event) error
	UpdateEvent(context.Context, storage.Event) error
	DeleteEvent(context.Context, storage.Event) error
	GetEvents(context.Context) ([]storage.Event, error)
}

func New(logger logger.Logger, storage Storage) *App {
	return &App{
		logger,
		storage,
	}
}

func (a *App) Index(ctx context.Context) map[string]string {
	return map[string]string{"message": "hello world"}
}

func (a *App) CreateEvent(ctx context.Context, event storage.Event) error {
	if err := a.storage.CreateEvent(ctx, event); err != nil {
		a.logger.Error("cannot create event error")
		return fmt.Errorf("cannot create event: %w", err)
	}

	return nil
}

func (a *App) UpdateEvent(ctx context.Context, e storage.Event) error {
	if err := a.storage.UpdateEvent(ctx, e); err != nil {
		a.logger.Error("cannot update event error")
		return fmt.Errorf("cannot update event: %w", err)
	}
	return nil
}

func (a *App) DeleteEvent(ctx context.Context, e storage.Event) error {
	if err := a.storage.DeleteEvent(ctx, e); err != nil {
		a.logger.Error("cannot delete event error")
		return fmt.Errorf("cannot delete event: %w", err)
	}
	return nil
}

func (a *App) GetEvents(ctx context.Context) ([]storage.Event, error) {
	events, err := a.storage.GetEvents(ctx)
	if err != nil {
		a.logger.Error("cannot get events error")
		return nil, fmt.Errorf("cannot get events: %w", err)
	}

	return events, nil
}
