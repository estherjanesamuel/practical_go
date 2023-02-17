package models


import (
	"github.com/lib/pq"
  )

type Recipe struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Ingredients  pq.StringArray `gorm:"type:text[]"`
	IsHalal      bool     `json:"isHalal"`
	IsVegetarian bool     `json:"isVegetarian"`
	Description  string   `json:"description,omitempty"`
	Rating       float64  `json:"rating"`
}
