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
