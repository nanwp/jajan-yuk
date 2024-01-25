package entity

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       int64     `json:"price"`
	Category    Category  `json:"category"`
	Variant     []Variant `json:"variant"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Variant struct {
	ID               int64         `json:"id"`
	Name             string        `json:"name"`
	TotalPrice       int64         `json:"total_price"`
	CountVariantType int64         `json:"count_variant_type"`
	VariantTypes     []VariantType `json:"variant_types,omitempty"`
}

type VariantType struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	VariantID int64  `json:"variant_id"`
}
