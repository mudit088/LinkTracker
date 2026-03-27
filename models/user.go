package models

type User struct {
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}