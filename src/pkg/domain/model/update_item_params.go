package domain

type UpdateItemParams struct {
	Name        string
	Variant     string
	Description string
	BuyValue    int
	SellValue   int
	Weight      int
	Durability  int
	MaxStack    int
	Type        ItemType
	Shape       Shape
}
