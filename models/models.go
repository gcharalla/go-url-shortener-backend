package models

import "gorm.io/gorm"

type Goly struct {
	gorm.Model
	ID       uint64 `json:"id" gorm:"primary_key"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}
