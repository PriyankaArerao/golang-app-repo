package models

type Coach struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
