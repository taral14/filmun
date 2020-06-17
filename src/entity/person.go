package entity

type Person struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	OriginalName string `json:"original_name"`
	KpInternalId int    `json:"-"`
	ImageUrl     string `json:"image"`
}
