package pedagang

import "github.com/nanwp/jajan-yuk/user/core/entity"

type CreatePedagangResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    Pedagang `json:"data"`
}

type Pedagang struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id"`
	Image        string  `json:"image"`
	NameMerchant string  `json:"name_merchant"`
	Phone        string  `json:"phone"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	IsActive     bool    `json:"is_active"`
	Distance     float64 `json:"distance"`
}

func (p *Pedagang) ToEntity() entity.Pedagang {
	return entity.Pedagang{
		ID:           p.ID,
		UserID:       p.UserID,
		Image:        p.Image,
		NameMerchant: p.NameMerchant,
		Phone:        p.Phone,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		IsActive:     p.IsActive,
		Distance:     p.Distance,
	}
}

// {
//     "success": true,
//     "message": "Success create pedagang",
//     "data": {
//         "ID": "9b0287b1-8baf-4724-ae7f-6ee881fa13b5",
//         "user_id": "29231d5a-cdd9-4014-a837-898735768641",
//         "image": "20240124001626.png",
//         "name_merchant": "baso cuanky asoy",
//         "phone": "08213213",
//         "latitude": 0,
//         "longitude": 0,
//         "is_active": true,
//         "distance": 0
//     }
// }
