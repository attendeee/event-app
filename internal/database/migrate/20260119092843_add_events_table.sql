-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    owner_id INTEGER,

    name TEXT NOT NULL,
    description TEXT NOT NULL,

    date TEXT NOT NULL,
    location TEXT NOT NULL,

    FOREIGN KEY (owner_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
