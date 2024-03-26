// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: item.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createItem = `-- name: CreateItem :one
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
  id, name, description, variant, buy_value, sell_value, weight, durability, max_stack, height, width, rawshape, created_at, type
`

type CreateItemParams struct {
	Name        pgtype.Text
	Description pgtype.Text
	Variant     pgtype.Text
	BuyValue    pgtype.Int4
	SellValue   pgtype.Int4
	Weight      pgtype.Int4
	Durability  pgtype.Int4
	MaxStack    pgtype.Int4
	Height      pgtype.Int4
	Width       pgtype.Int4
	Rawshape    pgtype.Text
	Type        NullItemType
	CreatedAt   pgtype.Timestamp
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRow(ctx, createItem,
		arg.Name,
		arg.Description,
		arg.Variant,
		arg.BuyValue,
		arg.SellValue,
		arg.Weight,
		arg.Durability,
		arg.MaxStack,
		arg.Height,
		arg.Width,
		arg.Rawshape,
		arg.Type,
		arg.CreatedAt,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Variant,
		&i.BuyValue,
		&i.SellValue,
		&i.Weight,
		&i.Durability,
		&i.MaxStack,
		&i.Height,
		&i.Width,
		&i.Rawshape,
		&i.CreatedAt,
		&i.Type,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :one
DELETE FROM item
WHERE
  id = $1
RETURNING
  id, name, description, variant, buy_value, sell_value, weight, durability, max_stack, height, width, rawshape, created_at, type
`

func (q *Queries) DeleteItem(ctx context.Context, id int32) (Item, error) {
	row := q.db.QueryRow(ctx, deleteItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Variant,
		&i.BuyValue,
		&i.SellValue,
		&i.Weight,
		&i.Durability,
		&i.MaxStack,
		&i.Height,
		&i.Width,
		&i.Rawshape,
		&i.CreatedAt,
		&i.Type,
	)
	return i, err
}

const getItem = `-- name: GetItem :one
SELECT
  id, name, description, variant, buy_value, sell_value, weight, durability, max_stack, height, width, rawshape, created_at, type
FROM
  item
WHERE
  id = $1
`

func (q *Queries) GetItem(ctx context.Context, id int32) (Item, error) {
	row := q.db.QueryRow(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Variant,
		&i.BuyValue,
		&i.SellValue,
		&i.Weight,
		&i.Durability,
		&i.MaxStack,
		&i.Height,
		&i.Width,
		&i.Rawshape,
		&i.CreatedAt,
		&i.Type,
	)
	return i, err
}

const getItemByType = `-- name: GetItemByType :many
SELECT
  id, name, description, variant, buy_value, sell_value, weight, durability, max_stack, height, width, rawshape, created_at, type
FROM
  item
WHERE
  type = $1
ORDER BY
  id
`

func (q *Queries) GetItemByType(ctx context.Context, type_ NullItemType) ([]Item, error) {
	rows, err := q.db.Query(ctx, getItemByType, type_)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Variant,
			&i.BuyValue,
			&i.SellValue,
			&i.Weight,
			&i.Durability,
			&i.MaxStack,
			&i.Height,
			&i.Width,
			&i.Rawshape,
			&i.CreatedAt,
			&i.Type,
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

const listItems = `-- name: ListItems :many
SELECT
  id, name, description, variant, buy_value, sell_value, weight, durability, max_stack, height, width, rawshape, created_at, type
FROM
  item
ORDER BY
  id
`

func (q *Queries) ListItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.Query(ctx, listItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Variant,
			&i.BuyValue,
			&i.SellValue,
			&i.Weight,
			&i.Durability,
			&i.MaxStack,
			&i.Height,
			&i.Width,
			&i.Rawshape,
			&i.CreatedAt,
			&i.Type,
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

const updateItem = `-- name: UpdateItem :one
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
  id, name, description, variant, buy_value, sell_value, weight, durability, max_stack, height, width, rawshape, created_at, type
`

type UpdateItemParams struct {
	Name        pgtype.Text
	Description pgtype.Text
	Variant     pgtype.Text
	BuyValue    pgtype.Int4
	SellValue   pgtype.Int4
	Weight      pgtype.Int4
	Durability  pgtype.Int4
	MaxStack    pgtype.Int4
	Height      pgtype.Int4
	Width       pgtype.Int4
	Rawshape    pgtype.Text
	Type        NullItemType
	CreatedAt   pgtype.Timestamp
	ID          int32
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRow(ctx, updateItem,
		arg.Name,
		arg.Description,
		arg.Variant,
		arg.BuyValue,
		arg.SellValue,
		arg.Weight,
		arg.Durability,
		arg.MaxStack,
		arg.Height,
		arg.Width,
		arg.Rawshape,
		arg.Type,
		arg.CreatedAt,
		arg.ID,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Variant,
		&i.BuyValue,
		&i.SellValue,
		&i.Weight,
		&i.Durability,
		&i.MaxStack,
		&i.Height,
		&i.Width,
		&i.Rawshape,
		&i.CreatedAt,
		&i.Type,
	)
	return i, err
}