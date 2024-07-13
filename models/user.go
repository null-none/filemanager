package models

type User struct {
  ID         uint   `json:"id"`
  Phone      string `json:"phone"`
  Password   string `json:"password"`
}