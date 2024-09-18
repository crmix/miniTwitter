package entities

import (
	"time"
)

type RegistrReq struct {
	ID           string
	Username     string `json:"username" gorm:"column:username"`
	Name         string `json:"name" gorm:"column:name"`
	Password     string `json:"password" gorm:"-"`
	HashPassword string `json:"-" gorm:"column:password"`
}

type RegistrRes struct {
	ID    string `json:"id"`
	Token Tokens `json:"tokens"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	Username string `json:"username"`
	Token    Tokens `json:"token"`
}

type PassId struct {
	ID       string `json:"id"`
	Password string `json:"password" binding:"required"`
}

type BlockedToken struct {
	ID        string `gorm:"primaryKey"`
	Token     string `json:"token" gorm:"unique;not null"`
	ExpiresAt int64  `gorm:"not null"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	ID           string   `gorm:"primarykey"`
	Username     string   `gorm:"unique;not null" json:"username"`
	Name         string   `json:"name"`
	Bio          string   `json:"bio"`
	ProfileImage string   `json:"profile_image"`
	Followers    []Follow `gorm:"foreignKey:FollowedID" json:"followers"`
	Following    []Follow `gorm:"foreignKey:FollowerID" json:"following"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Follow struct {
	ID         string
	FollowerID string `json:"follower_id"`
	FollowedID string `json:"followed_id"`
	CreatedAt  time.Time
}

type Like struct {
	ID        string
	TweetID   string `json:"tweet_id"`
	UserId    string `json:"user_id"`
	CreatedAt time.Time
}

type Retweet struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	State           string `json:"state"`
	OriginalTweetID string `json:"original_tweet_id"`
	CreatedAt       time.Time
}

type Tweet struct {
	ID        string
	Content   string    `gorm:"not null" json:"content"`
	ImageURL  string    `json:"image_url"`
	VideoURL  string    `json:"video_url"`
	UserName  string    `json:"user_name"`
	State     string    `json:"state"`
	User      User      `json:"user"`
	Likes     []Like    `json:"likes"`
	Retweets  []Retweet `json:"retweets"`
	CreateAt  time.Time
	UpdatedAt time.Time
}

type CreateTweetReq struct {
	ID       string
	UserId   string `json:"-" gorm:"column:user_id"`
	Content  string `json:"content" binding:"required"`
	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`
	State    string `json:"state"`
	CreatedAt time.Time
}
