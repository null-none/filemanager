package models

type User struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Phone  string `json:"phone"`
  Year   string `json:"year"`
}