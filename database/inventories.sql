-- name: ListInventories :many
SELECT
  *
FROM
  inventory
ORDER BY
  id;

-- name: GetInventory :one
SELECT
  *
FROM
  inventory
WHERE
  id = $1;

-- name: CreateInventory :one
INSERT INTO
  inventory (
    user_id,
    invname,
    width,
    height,
    max_weight,
    created_at
  )
VALUES
  ($1, $2, $3, $4, $5, $6)
RETURNING
  *;

-- name: UpdateInventory :one
UPDATE inventory
SET
  user_id = $1,
  invname = $2,
  width = $3,
  height = $4,
  max_weight = $5,
  created_at = $6
WHERE
  id = $7
RETURNING
  *;

-- name: DeleteInventory :one
DELETE FROM inventory
WHERE
  id = $1
RETURNING
  *;
