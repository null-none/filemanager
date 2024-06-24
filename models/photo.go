package models

type Photo struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Path string `json:"path"`
	Tour int    `json:"tour"`
}
