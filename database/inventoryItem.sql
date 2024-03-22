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
