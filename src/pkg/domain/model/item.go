package domain

type Item struct {
	ItemMeta   ItemMeta
	Name       string
	Text       string
	Type       ItemType
	Durability int
	SellValue  int
	BuyValue   int
	Variant    string
}
