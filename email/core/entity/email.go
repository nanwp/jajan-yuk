package entity

type Email struct {
	Title    string `json:"sender,omitempty"`
	Receiver string `json:"recive,omitempty"`
	Subject  string `json:"subject,omitempty"`
	Body     string `json:"body,omitempty"`
}
