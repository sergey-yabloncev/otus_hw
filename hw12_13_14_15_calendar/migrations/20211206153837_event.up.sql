CREATE TABLE events
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    event_at    TIMESTAMP    NOT NULL,
    start_at    TIMESTAMP    NOT NULL,
    end_at      TIMESTAMP    NOT NULL,
    notify_at   TIMESTAMP    NOT NULL,
    is_notify   boolean      NOT NULL DEFAULT false
--     user_id     integer REFERENCES users ON DELETE CASCADE
);
