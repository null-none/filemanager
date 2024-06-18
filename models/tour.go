package models

type Tour struct {
  ID            uint   `json:"id" gorm:"primary_key"`
  Day           string `json:"day"`
  Title         string `json:"title"`
  Description   string `json:"description"`
}