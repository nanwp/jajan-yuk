package entity

type Email struct {
	Title    string `json:"title"`
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}
