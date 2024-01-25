package product_client

import "github.com/nanwp/jajan-yuk/pedagang/core/entity"

type GetPedagangByUserIDResponse struct {
	Succeess bool      `json:"success"`
	Message  string    `json:"message"`
	Data     []Product `json:"data,omitempty"`
}

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

func (v *Variant) ToEntity() entity.Variant {
	variantTypes := make([]entity.VariantType, 0)

	for _, variantType := range v.VariantTypes {
		variantTypes = append(variantTypes, entity.VariantType{
			ID:        variantType.ID,
			Name:      variantType.Name,
			Price:     variantType.Price,
			VariantID: variantType.VariantID,
		})
	}

	return entity.Variant{
		ID:               v.ID,
		Name:             v.Name,
		TotalPrice:       v.TotalPrice,
		CountVariantType: v.CountVariantType,
		VariantTypes:     variantTypes,
	}
}

func (p *Product) ToEntity() entity.Product {
	variants := make([]entity.Variant, 0)

	for _, variant := range p.Variant {
		variants = append(variants, variant.ToEntity())
	}

	return entity.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
		Price:       p.Price,
		Category: entity.Category{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		},
		Variant: variants,
	}
}
