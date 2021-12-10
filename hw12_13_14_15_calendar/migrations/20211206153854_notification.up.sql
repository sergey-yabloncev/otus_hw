CREATE TABLE notifications
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    event_at    TIMESTAMP    NOT NULL,
    user_id     integer REFERENCES users ON DELETE CASCADE,
    event_id    integer REFERENCES events ON DELETE CASCADE
);
