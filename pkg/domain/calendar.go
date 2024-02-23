package domain

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name       string
	Date       string
	CalendarID uint
	CustomerID uint
	Calendar   Calendar `gorm:"foreignKey:CalendarID"`
	Customer   Customer `gorm:"foreignKey:CustomerID"`
}

type Calendar struct {
	gorm.Model
	Name   string  `gorm:"unique" json:"name"`
	Events []Event `gorm:"foreignKey:CalendarID"`
}
