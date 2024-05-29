// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package repo

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type ItemType string

const (
	ItemTypeConsumable       ItemType = "consumable"
	ItemTypeArmor            ItemType = "armor"
	ItemTypeRangedWeapon     ItemType = "rangedWeapon"
	ItemTypeMeleeWeapon      ItemType = "meleeWeapon"
	ItemTypeConsumableWeapon ItemType = "consumableWeapon"
	ItemTypeQuest            ItemType = "quest"
	ItemTypeResource         ItemType = "resource"
)

func (e *ItemType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ItemType(s)
	case string:
		*e = ItemType(s)
	default:
		return fmt.Errorf("unsupported scan type for ItemType: %T", src)
	}
	return nil
}

type NullItemType struct {
	ItemType ItemType
	Valid    bool // Valid is true if ItemType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullItemType) Scan(value interface{}) error {
	if value == nil {
		ns.ItemType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ItemType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullItemType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ItemType), nil
}

type Inventory struct {
	ID        int32
	Invname   pgtype.Text
	UserID    pgtype.Int4
	Width     pgtype.Int4
	Height    pgtype.Int4
	MaxWeight pgtype.Int4
	CreatedAt pgtype.Timestamp
}

type InventoryItem struct {
	ID             int32
	InventoryID    pgtype.Int4
	ItemID         pgtype.Int4
	Quantity       pgtype.Int4
	PositionX      pgtype.Int4
	PositionY      pgtype.Int4
	Rotation       pgtype.Int4
	DurabilityLeft pgtype.Int4
	CreatedAt      pgtype.Timestamp
}

type Item struct {
	ID         int32
	Name       pgtype.Text
	Text       pgtype.Text
	Variant    pgtype.Text
	BuyValue   pgtype.Int4
	SellValue  pgtype.Int4
	Weight     pgtype.Int4
	Durability pgtype.Int4
	MaxStack   pgtype.Int4
	Height     pgtype.Int4
	Width      pgtype.Int4
	Rawshape   pgtype.Text
	CreatedAt  pgtype.Timestamp
	Type       NullItemType
}

type User struct {
	ID        int32
	Username  pgtype.Text
	Email     pgtype.Text
	CreatedAt pgtype.Timestamp
}
