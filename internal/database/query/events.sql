-- name: GetAllEvents :many
SELECT * FROM events;

-- name: GetEventByOwner :many
SELECT * FROM events 
WHERE owner_id = (?);

-- name: GetEventByName :many
SELECT * FROM events 
WHERE name = (?);

-- name: CreateEvent :exec
INSERT INTO 
events(owner_id, name, description, date, location) 
VALUES (?, ?, ?, ?, ?);
