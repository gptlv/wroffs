package main

type Record struct {
	ID             string `csv:"id"`
	ISC            string `csv:"isc"`
	Flaw           string `csv:"flaw"`
	Decision       string `csv:"decision"`
	Serial         string `csv:"serial"`
	Name           string `csv:"name"`
	InventoryID    string `csv:"inventory_id"`
	Date           string `csv:"date"`
	DepartmentLead string `csv:"department_lead"`
	TeamLead       string `csv:"team_lead"`
	Director       string `csv:"director"`
}
