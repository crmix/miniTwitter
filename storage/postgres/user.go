package postgres

import (
	"context"
	"errors"
	"fmt"
	"miniTwitter/constants"
	"miniTwitter/entities"
	e "miniTwitter/errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdmin(db *gorm.DB) *adminRepo {
	return &adminRepo{db: db}
}

func (a adminRepo) Registration(ctx context.Context, req entities.RegistrReq) error {
	res := a.db.WithContext(ctx).Table("users").Create(&req)
	if res.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(res.Error, &pgErr) && pgErr.Code == constants.PGUniqueKeyViolationCode {
			return e.ErrAccountAlreadyExists
		}
		return fmt.Errorf("error in Registration: %w", res.Error)
	}
	return nil
}

func (a adminRepo) Login(ctx context.Context, userName string) (entities.PassId, error) {
	var password entities.PassId
	err := a.db.WithContext(ctx).Table("users").Select("id", "password").Where("username=?", userName).First(&password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.PassId{}, errors.New("user not found with this username")
		}
		return entities.PassId{}, err
	}
	return password, nil
}

func (r adminRepo) Logout(ctx context.Context, token entities.BlockedToken) error {

	err := r.db.WithContext(ctx).Table("blocked_tokens").Create(&token).Error
	if err != nil {
		return fmt.Errorf("error in Logout: %w", err)
	}
	return nil
}

func (r adminRepo) CreateTweet(ctx context.Context, tweet entities.CreateTweetReq) error {

	err := r.db.WithContext(ctx).Table("tweets").Create(&tweet).Error
	if err != nil {
		return fmt.Errorf("error in CreateTweeet: %w", err)
	}
	return nil
}

func (r adminRepo) Follow(ctx context.Context, req entities.Follow) error {
	err := r.db.WithContext(ctx).Table("follows").Create(&req).Error
	if err != nil {
		return fmt.Errorf("error in Follow: %w", err)
	}
	return nil
}

func (r adminRepo) Unfollow(ctx context.Context, req entities.Follow) error {
	err := r.db.WithContext(ctx).Table("follows").Where("follower_id = ? AND followed_id = ?", req.FollowerID, req.FollowedID).Delete(nil).Error
	if err != nil {
		return fmt.Errorf("error in Follow: %w", err)
	}
	return nil
}

func (r adminRepo) GetFollowers(ctx context.Context, userId string) ([]entities.User, error) {
	var followers []entities.User
	err := r.db.Table("follows").
		Select("users.*").
		Joins("JOIN users ON users.id = follows.follower_id").
		Where("follows.followed_id = ?", userId).
		Scan(&followers).Error
	if err != nil {
		return []entities.User{}, fmt.Errorf("error in Follow: %w", err)
	}
	return followers, nil
}

func (r adminRepo) GetFollowings(ctx context.Context, userId string) ([]entities.User, error) {
	var followings []entities.User
	err := r.db.Table("follows").
        Select("users.*").
        Joins("JOIN users ON users.id = follows.followed_id").
        Where("follows.follower_id = ?", userId).
        Scan(&followings).Error
	if err != nil {
		return []entities.User{}, fmt.Errorf("error in Follow: %w", err)
	}
	return followings, nil
}

func (r adminRepo) LikeTweet(ctx context.Context, req entities.Like) error {
	err := r.db.WithContext(ctx).Table("likes").Create(&req).Error
	if err != nil {
		return fmt.Errorf("error in LikeTweet: %w", err)
	}
	return nil
}

func (r adminRepo) UnlikeTweet(ctx context.Context, req entities.Like) error {
	err := r.db.WithContext(ctx).Table("likes").
	Where("tweet_id = ? AND user_id = ?", req.TweetID, req.UserId).
	Delete(nil).Error
	if err != nil {
		return fmt.Errorf("error in LikeTweet: %v", err)
	}
	return nil
}

func (r adminRepo) Retweet(ctx context.Context, req entities.Retweet) error {
	err := r.db.WithContext(ctx).Table("tweets").Create(&req).Error
	if err != nil {
		return fmt.Errorf("error in Retweet: %v", err)
	}
	return nil
}

func (r adminRepo) Search(ctx context.Context, req string) ([]entities.User, error){
	var users []entities.User
	err := r.db.Table("users").
		Where("username ILIKE ? OR name ILIKE ?", "%"+req+"%", "%"+req+"%").
		Find(&users).Error
		if err != nil {
			return []entities.User{}, fmt.Errorf("error in Retweet: %v", err)
		}
		return users, nil
}