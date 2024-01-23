package auth

type GetCurrentUserResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
	Data    CurrentUserData `json:"data"`
}

type CurrentUserData struct {
	User User `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     Role   `json:"role"`
}

type ValidateSecretKeyResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
	Data    SecretKey `json:"data"`
}

type SecretKey struct {
	ID     string `json:"id"`
	Serial string `json:"serial"`
	Name   string `json:"name"`
	Role   Role   `json:"role"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}
