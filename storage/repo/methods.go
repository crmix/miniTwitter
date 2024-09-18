package repo

import (
	"context"
	"miniTwitter/entities"
)

// IAdminStorage account storage interface
type IAdminStorage interface {
	Registration(ctx context.Context, req entities.RegistrReq) error
	Login(ctx context.Context, userName string)(entities.PassId, error)
	Logout(ctx context.Context, req entities.BlockedToken) error
	CreateTweet(ctx context.Context, req entities.CreateTweetReq) error
	Follow(ctx context.Context, req entities.Follow) error
	Unfollow(ctx context.Context, req entities.Follow) error
	GetFollowers(ctx context.Context, userId string) ([]entities.User, error) 
	GetFollowings(ctx context.Context, userId string) ([]entities.User, error)
	LikeTweet(ctx context.Context, req entities.Like) error
	UnlikeTweet(ctx context.Context, req entities.Like) error 
	Retweet(ctx context.Context, req entities.Retweet) error 
	Search(ctx context.Context, req string) ([]entities.User, error)
}