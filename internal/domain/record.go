package domain

type Record struct {
	ID          string `csv:"id"`
	ISC         string `csv:"isc"`
	Flaw        string `csv:"flaw"`
	Decision    string `csv:"decision"`
	Serial      string `csv:"serial"`
	Name        string `csv:"name"`
	InventoryID string `csv:"inventory_id"`
	Date        string `csv:"date"`
	Boss        string `csv:"boss"`
	Lead        string `csv:"lead"`
}
