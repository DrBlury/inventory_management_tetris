package domain

type CreateItemParams struct {
	Name       string
	Variant    string
	Text       string
	BuyValue   int
	SellValue  int
	Weight     int
	Durability int
	MaxStack   int
	Type       ItemType
	Shape      Shape
}
