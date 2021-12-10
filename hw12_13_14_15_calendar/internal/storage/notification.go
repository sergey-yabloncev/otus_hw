package storage

import "time"

/*Notification
ID события;
Заголовок события;
Дата события;
Пользователь, которому отправлять.
*/
type Notification struct {
	ID      uint64
	Title   string
	EventAt time.Time
	Event   Event
	User    User
}
