package types

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint      `gorm:"primary"`
	HashedPassword string    `gorm:"size:50;not null"`
	Salt           string    `gorm:"size:20;not null"`
	FirstName      string    `gorm:"size:50;not null"`
	LastName       string    `gorm:"size:50;not null"`
	DateOfBirth    time.Time `gorm:"not null"`
	Email          string    `gorm:"size:50;not null"`
	UserName       string    `gorm:"size:50;not null;index:idx_username"`
	Following      []*User   `gorm:"many2many:following;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:friend_id"`
	Follower       []*User   `gorm:"many2many:following;foreignKey:id;joinForeignKey:friend_id;References:id;joinReferences:user_id"`
	Posts          []*Post   `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "user"
}

type Post struct {
	gorm.Model
	ContentText      string     `gorm:"size:500"`
	ContentImagePath string     `gorm:"size:255"`
	UserID           uint       `gorm:"not null"`
	Visible          bool       `gorm:"not null"`
	Comments         []*Comment `gorm:"foreignKey:PostID"`
	LikedUsers       []*User    `gorm:"many2many:like;foreignKey:id;joinForeignKey:post_id;References:id;joinReferences:user_id"`
}

func (Post) TableName() string {
	return "post"
}

type Comment struct {
	gorm.Model
	Content string `gorm:"size:255;not null"`
	UserID  uint   `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
}

func (Comment) TableName() string {
	return "comment"
}

type Like struct {
	gorm.Model
	UserID uint `gorm:"primaryKey"`
	PostID uint `gorm:"primaryKey"`
}

func (Like) TableName() string {
	return "like"
}
