package userLogin

import "time"

type RequestResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserDetails struct {
	Id           int        `json:"id" gorm:"primaryKey"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	CreatedAt    *time.Time `json:"created_at"`
	LastLoggedIn *time.Time `json:"last_logged_in"`
}

type UserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
