-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS attendees (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    event_id INTEGER,

    FOREIGN KEY (user_id) REFERENCES users(id)
    FOREIGN KEY (event_id) REFERENCES events(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attendees;
-- +goose StatementEnd
