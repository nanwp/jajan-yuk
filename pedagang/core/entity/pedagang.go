package entity

type Pedagang struct {
	ID           string  `json:"ID"`
	UserID       string  `json:"user_id"`
	Image        string  `json:"image"`
	NameMerchant string  `json:"name_merchant"`
	Phone        string  `json:"phone"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	IsActive     bool    `json:"is_active"`
	Distance     float64 `json:"distance"`
}

type UpdateLocationRequest struct {
	UserID    string  `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GetAllPedagangNearbyRequest struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MaxDistance float64 `json:"max_distance"`
	Keyword     string  `json:"keyword"`
}
