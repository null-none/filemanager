package models

type Folder struct {
  ID            uint   `json:"id" gorm:"primary_key"`
  Prefix        string `json:"prefix"`
  Title         string `json:"title"`
  Description   string `json:"description"`
}