-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = (?);

-- name: CreateUser :one
INSERT INTO 
users(name, email, password) 
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateUserInfo :exec
UPDATE users 
SET 
    name = COALESCE(NULLIF(?, ''), name),
    email = COALESCE(NULLIF(?, ''), email)

WHERE id = (?);

-- name: UpdateUserPassword :exec
UPDATE users 
SET password = (?) 
WHERE id = (?);

-- name: DeleteUserById :exec
DELETE FROM users 
WHERE id = (?);

