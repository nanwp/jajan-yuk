package entity

type Product struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Price      int64     `json:"price"`
	Category   Category  `json:"category"`
	Variant    []Variant `json:"variant"`
	PedagangID string    `json:"pedagang_id"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Variant struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	VariantTypes []VariantType `json:"variant_types"`
}

type VariantType struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	VariantID int64  `json:"variant_id"`
}
