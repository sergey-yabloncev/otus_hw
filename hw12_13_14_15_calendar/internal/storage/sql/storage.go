package sqlstorage

import (
	"context"

	_ "github.com/jackc/pgx/v4/stdlib" //nolint
	"github.com/jmoiron/sqlx"
	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	db *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) error {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return err
	}

	s.db = db

	return s.db.Ping()
}

func (s *Storage) Close(_ context.Context) error {
	return s.db.Close()
}

func (s *Storage) CreateEvent(ctx context.Context, e storage.Event) error {
	_, err := s.db.NamedExecContext(ctx, `INSERT INTO events (
                    id, 
                    title, 
                    description,
                    event_at,
                    start_at,
                    end_at,
                    notify_at,
                    is_notify
--                     user_id
  		  ) VALUES (
	     			:id, 
                    :title, 
                    :description,
                    :event_at,
                    :start_at,
                    :end_at,
                    :notify_at,
                    :is_notify
--                     :user_id
)`, e)

	return err
}

func (s *Storage) UpdateEvent(ctx context.Context, e storage.Event) error {
	_, err := s.db.ExecContext(ctx, `UPDATE events  
			SET title = $2, 
				description = $3,
				event_at = $4,
				start_at = $5,
				end_at = $6,
				notify_at = $7,
				is_notify = $8
			WHERE id = $1`,
		e.ID, e.Title, e.Description, e.EventAt, e.StartAt, e.StartAt, e.StartAt, e.NotifyAt, e.IsNotify)

	return err
}

func (s *Storage) DeleteEvent(ctx context.Context, e storage.Event) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM events WHERE id = $1", e.ID)

	return err
}

func (s *Storage) GetEvents(ctx context.Context) ([]storage.Event, error) {
	event := storage.Event{}
	events := []storage.Event{}

	prepare, err := s.db.PrepareNamed(`SELECT * FROM events`)
	if err != nil {
		return nil, err
	}
	err = prepare.Select(&events, event)
	if err != nil {
		return nil, err
	}

	return events, nil
}
