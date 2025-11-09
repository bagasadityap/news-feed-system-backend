package models

import "time"

type Follow struct {
	FollowerID uint      `gorm:"not null" json:"follower_id"`
	FolloweeID uint      `gorm:"not null" json:"followee_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`

	Follower User `gorm:"foreignKey:FollowerID;references:ID;constraint:OnDelete:CASCADE"`
	Followee User `gorm:"foreignKey:FolloweeID;references:ID;constraint:OnDelete:CASCADE"`
}