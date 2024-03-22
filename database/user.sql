-- name: GetUser :one
SELECT
  *
FROM
  "user"
WHERE
  id = $1;

-- name: GetUserByUsername :one
SELECT
  *
FROM
  "user"
WHERE
  username = $1;

-- name: ListUsers :many
SELECT
  *
FROM
  "user"
ORDER BY
  id;

-- name: CreateUser :one
INSERT INTO
  "user" (username, email, salt, password_hash, created_at)
VALUES
  ($1, $2, $3, $4, $5)
RETURNING
  *;

-- name: UpdateUser :one
UPDATE "user"
SET
  username = $1,
  email = $2,
  salt = $3,
  password_hash = $4,
  created_at = $5
WHERE
  id = $6
RETURNING
  *;

-- name: DeleteUser :one
DELETE FROM "user"
WHERE
  id = $1
RETURNING
  *;
