package models

type File struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Folder int    `json:"folder"`
}
