package entity

type Payment struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Status  string `json:"status"`
	Ammount int64  `json:"ammount"`
	SnapURL string `json:"snap_url"`
}

type PaymentRequest struct {
	UserID  string `json:"user_id"`
	Ammount int64  `json:"ammount"`
}

type PaymentResponse struct {
	SnapURL string `json:"snap_url"`
}
