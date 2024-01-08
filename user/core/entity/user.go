package entity

type User struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Username    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Address     string `json:"address,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Role        Role   `json:"role"`
}

type Role struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
