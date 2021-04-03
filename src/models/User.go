package models

import (
	"time"

	"gorm.io/gorm"
)

// User struct to model database table
type User struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey; autoIncrement; not null" json:"id"`
	Name      string    `gorm:"not null" json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `gorm:"unique; not null" json:"email,omitempty"`
	Password  string    `gorm:"not null" json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
