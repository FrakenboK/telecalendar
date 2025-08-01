package models

import (
	"time"

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
	Name string `gorm:"unique;not null" json:"name,omitempty"`

	Users  []User  `gorm:"many2many:users_calendars"`
	Events []Event `gorm:"many2many:calendars_events"`
}

type Event struct {
	ID   uint      `gorm:"primarykey" json:"-"`
	Name string    `gorm:"not null" json:"name,omitempty"`
	Type EventType `gorm:"type:varchar(100);not null" json:"type,omitempty"`

	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`

	Hour   int `json:"hour,omitempty"`
	Minute int `json:"minute,omitempty"`

	Aggressiveness AggressivenessType `gorm:"type:varchar(100);not null" json:"aggressiveness,omitempty"`

	Calendars []Calendar `gorm:"many2many:calendars_events,omitempty" json:"-"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
