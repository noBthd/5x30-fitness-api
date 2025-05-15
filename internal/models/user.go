package models

type User struct {
	ID		 string `json:"id"`
	Username string `json:"username"`
	Email	 string `json:"email"`
	Password string `json:"password"`
	RegDate  string `json:"registration_date"`
}

