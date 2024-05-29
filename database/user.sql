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
  "user" (username, email, created_at)
VALUES
  ($1, $2, $3)
RETURNING
  *;

-- name: UpdateUser :one
UPDATE "user"
SET
  username = $1,
  email = $2,
  created_at = $3
WHERE
  id = $4
RETURNING
  *;

-- name: DeleteUser :one
DELETE FROM "user"
WHERE
  id = $1
RETURNING
  *;
