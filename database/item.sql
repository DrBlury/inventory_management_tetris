-- name: GetItem :one
SELECT
  *
FROM
  item
WHERE
  id = $1;

-- name: ListItems :many
SELECT
  *
FROM
  item
ORDER BY
  id;

-- name: CreateItem :one
INSERT INTO
  item (
    name,
    text,
    variant,
    buy_value,
    sell_value,
    weight,
    durability,
    max_stack,
    height,
    width,
    rawshape,
    type,
    created_at
  )
VALUES
  (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13
  )
RETURNING
  *;

-- name: UpdateItem :one
UPDATE "item"
SET
  name = $1,
  text = $2,
  variant = $3,
  buy_value = $4,
  sell_value = $5,
  weight = $6,
  durability = $7,
  max_stack = $8,
  height = $9,
  width = $10,
  rawshape = $11,
  type = $12,
  created_at = $13
WHERE
  id = $14
RETURNING
  *;

-- name: DeleteItem :one
DELETE FROM item
WHERE
  id = $1
RETURNING
  *;

-- name: GetItemByType :many
SELECT
  *
FROM
  item
WHERE
  type = $1
ORDER BY
  id;
