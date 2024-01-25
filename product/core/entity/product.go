package entity

import "fmt"

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

func (c *Category) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("category name is required")
	}

	return nil
}

type Variant struct {
	ID               int64         `json:"id"`
	Name             string        `json:"name"`
	TotalPrice       int64         `json:"total_price"`
	CountVariantType int64         `json:"count_variant_type"`
	VariantTypes     []VariantType `json:"variant_types"`
}

type VariantType struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	VariantID int64  `json:"variant_id"`
}
