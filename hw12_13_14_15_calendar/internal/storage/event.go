package storage

import (
	"context"
	"time"
)

/*Event
ID - уникальный идентификатор события (можно воспользоваться UUID);
Заголовок - короткий текст;
Дата и время события;
Длительность события (или дата и время окончания);
Описание события - длинный текст, опционально;
ID пользователя, владельца события;
За сколько времени высылать уведомление, опционально.
*/
type Event struct {
	ID          string
	Title       string
	EventAt     time.Time
	StartAt     time.Time
	EndAt       time.Time
	Description string
	NotifyAt    time.Time
	IsNotify    time.Time
	// User        User
}

type EventRepo interface {
	CreateEvent(context.Context, Event) error
	UpdateEvent(context.Context, Event) error
	DeleteEvent(context.Context, Event) error
	GetEvents(context.Context) ([]Event, error)
}
