package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TelegramId int64 `gorm:"unique;not null"`
	Username   string
	Calendars  []Calendar `gorm:"many2many:users_calendars"`
}

type Calendar struct {
	gorm.Model
	Name string `gorm:"unique;not null"`

	Users  []User  `gorm:"many2many:users_calendars"`
	Events []Event `gorm:"many2many:calendars_events"`
}

type Event struct {
	gorm.Model
	Name string    `gorm:"not null"`
	Type EventType `gorm:"type:varchar(100);not null"`

	Calendars []Calendar `gorm:"many2many:calendars_events"`

	Year  int
	Month int
	Day   int

	Hour   int
	Minute int
}
