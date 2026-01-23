-- name: GetAllAttendeesForEvent :many
SELECT * FROM attendees
WHERE event_id = (?);

-- name: AddAttendee :exec
INSERT INTO 
attendees (event_id, user_id) 
VALUES (?, ?);

-- name: DeleteAttendee :exec
DELETE FROM attendees 
WHERE user_id = (?) AND event_id = (?);
