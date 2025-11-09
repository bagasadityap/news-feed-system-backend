package models

import "time"

type Post struct {
	ID		uint      `gorm:"primaryKey" json:"id"`
	UserID	uint      `gorm:"not null" json:"user_id"`
	Content	string    `gorm:"type:text;not null" json:"content"`
	CreatedAt	time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}