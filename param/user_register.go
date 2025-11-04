package param

import "time"

type RegisterRequest struct {
	FullName  string    `json:"full_name"`
	Handle    string    `json:"handle"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	AvatarURL string    `json:"avatar_url"`
}

type RegisterResponse struct {
	FullName  string    `json:"full_name"`
	Handle    string    `json:"handle"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
	AvatarURL string    `json:"avatar_url"`
}
