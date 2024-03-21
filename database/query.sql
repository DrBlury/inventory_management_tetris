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
    description,
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
  description = $2,
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

-- name: GetInventory :one
SELECT
  *
FROM
  inventory
WHERE
  id = $1;

-- name: ListInventories :many
SELECT
  *
FROM
  inventory
ORDER BY
  id;

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
  created_at = $5
WHERE
  id = $5
RETURNING
  *;

-- name: DeleteInventory :one
DELETE FROM inventory
WHERE
  id = $1
RETURNING
  *;

-- name: ListInventoryItems :many
SELECT
  *
FROM
  inventory_item
WHERE
  inventory_id = $1
ORDER BY
  id;

-- name: AddItemToInventory :one
INSERT INTO
  inventory_item (
    inventory_id,
    item_id,
    position_x,
    position_y,
    rotation,
    quantity,
    created_at
  )
VALUES
  ($1, $2, $3, $4, $5, $6, $7)
RETURNING
  *;

-- name: RemoveItemFromInventory :one
DELETE FROM inventory_item
WHERE
  id = $1
RETURNING
  *;

-- name: UpdateInventoryItem :one
UPDATE inventory_item
SET
  inventory_id = $1,
  item_id = $2,
  position_x = $3,
  position_y = $4,
  rotation = $5,
  quantity = $6,
  created_at = $7
WHERE
  id = $8
RETURNING
  *;
