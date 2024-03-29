// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: inventories.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createInventory = `-- name: CreateInventory :one
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
  id, invname, user_id, width, height, max_weight, created_at
`

type CreateInventoryParams struct {
	UserID    pgtype.Int4
	Invname   pgtype.Text
	Width     pgtype.Int4
	Height    pgtype.Int4
	MaxWeight pgtype.Int4
	CreatedAt pgtype.Timestamp
}

func (q *Queries) CreateInventory(ctx context.Context, arg CreateInventoryParams) (Inventory, error) {
	row := q.db.QueryRow(ctx, createInventory,
		arg.UserID,
		arg.Invname,
		arg.Width,
		arg.Height,
		arg.MaxWeight,
		arg.CreatedAt,
	)
	var i Inventory
	err := row.Scan(
		&i.ID,
		&i.Invname,
		&i.UserID,
		&i.Width,
		&i.Height,
		&i.MaxWeight,
		&i.CreatedAt,
	)
	return i, err
}

const deleteInventory = `-- name: DeleteInventory :one
DELETE FROM inventory
WHERE
  id = $1
RETURNING
  id, invname, user_id, width, height, max_weight, created_at
`

func (q *Queries) DeleteInventory(ctx context.Context, id int32) (Inventory, error) {
	row := q.db.QueryRow(ctx, deleteInventory, id)
	var i Inventory
	err := row.Scan(
		&i.ID,
		&i.Invname,
		&i.UserID,
		&i.Width,
		&i.Height,
		&i.MaxWeight,
		&i.CreatedAt,
	)
	return i, err
}

const getInventory = `-- name: GetInventory :one
SELECT
  id, invname, user_id, width, height, max_weight, created_at
FROM
  inventory
WHERE
  id = $1
`

func (q *Queries) GetInventory(ctx context.Context, id int32) (Inventory, error) {
	row := q.db.QueryRow(ctx, getInventory, id)
	var i Inventory
	err := row.Scan(
		&i.ID,
		&i.Invname,
		&i.UserID,
		&i.Width,
		&i.Height,
		&i.MaxWeight,
		&i.CreatedAt,
	)
	return i, err
}

const listInventories = `-- name: ListInventories :many
SELECT
  id, invname, user_id, width, height, max_weight, created_at
FROM
  inventory
ORDER BY
  id
`

func (q *Queries) ListInventories(ctx context.Context) ([]Inventory, error) {
	rows, err := q.db.Query(ctx, listInventories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Inventory
	for rows.Next() {
		var i Inventory
		if err := rows.Scan(
			&i.ID,
			&i.Invname,
			&i.UserID,
			&i.Width,
			&i.Height,
			&i.MaxWeight,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listInventoriesByUserID = `-- name: ListInventoriesByUserID :many
SELECT
  id, invname, user_id, width, height, max_weight, created_at
FROM
  inventory
WHERE
  user_id = $1
`

func (q *Queries) ListInventoriesByUserID(ctx context.Context, userID pgtype.Int4) ([]Inventory, error) {
	rows, err := q.db.Query(ctx, listInventoriesByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Inventory
	for rows.Next() {
		var i Inventory
		if err := rows.Scan(
			&i.ID,
			&i.Invname,
			&i.UserID,
			&i.Width,
			&i.Height,
			&i.MaxWeight,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInventory = `-- name: UpdateInventory :one
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
  id, invname, user_id, width, height, max_weight, created_at
`

type UpdateInventoryParams struct {
	UserID    pgtype.Int4
	Invname   pgtype.Text
	Width     pgtype.Int4
	Height    pgtype.Int4
	MaxWeight pgtype.Int4
	CreatedAt pgtype.Timestamp
	ID        int32
}

func (q *Queries) UpdateInventory(ctx context.Context, arg UpdateInventoryParams) (Inventory, error) {
	row := q.db.QueryRow(ctx, updateInventory,
		arg.UserID,
		arg.Invname,
		arg.Width,
		arg.Height,
		arg.MaxWeight,
		arg.CreatedAt,
		arg.ID,
	)
	var i Inventory
	err := row.Scan(
		&i.ID,
		&i.Invname,
		&i.UserID,
		&i.Width,
		&i.Height,
		&i.MaxWeight,
		&i.CreatedAt,
	)
	return i, err
}
