package domain

type Item struct {
	ID          int
	Name        string
	Description string
	Type		ItemType
	Durability  int
	MaxStack    int
	Weight      int
	Shape       Shape
	SellValue   int
	BuyValue	int
	Variant	 string
}

type Items struct {
	Items []*Item `yaml:"items"`
}
