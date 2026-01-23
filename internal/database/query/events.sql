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

-- name: UpdateEvent :exec
UPDATE events 
SET 
    description = COALESCE(NULLIF(?, ''), description),
    date = COALESCE(NULLIF(?, ''), date),
    location = COALESCE(NULLIF(?, ''), location)
WHERE id = (?);

-- name: DeleteEventById :exec
DELETE FROM events 
WHERE id = (?);

