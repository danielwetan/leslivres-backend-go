package models

type Users struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type Login struct {
	Username string `json:"username"`
}
