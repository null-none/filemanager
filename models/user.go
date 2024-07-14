package models

type User struct {
  ID         uint   `json:"id"`
  Password   string `json:"password"`
  Login      string `json:"login"`
}