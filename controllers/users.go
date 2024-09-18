package controllers

import (
	"context"
	"miniTwitter/configs"
	"miniTwitter/constants"
	"miniTwitter/entities"
	"miniTwitter/logger"
	pkgerrors "miniTwitter/pkg/errors"
	"miniTwitter/pkg/jwt"
	"miniTwitter/pkg/utils"
	"miniTwitter/storage"
	"net/http"

	"github.com/google/uuid"
)

type AdminController interface {
	Registration(ctx context.Context, req entities.RegistrReq) (entities.RegistrRes, error)
	Login(ctx context.Context, req entities.LoginReq) (entities.LoginRes, error)
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

type adminController struct {
	log     logger.LoggerI
	storage storage.Storage
	cfg     *configs.Configuration
}

func NewAdminController(log logger.LoggerI, storage storage.Storage) AdminController {
	return adminController{
		log:     log,
		storage: storage,
		cfg:     configs.Config(),
	}
}

func (a adminController) Registration(ctx context.Context, req entities.RegistrReq) (entities.RegistrRes, error) {
	a.log.Info("Registration started: ")
	var err error
	req.HashPassword, err = utils.HashPassword(req.Password)
	if err != nil {
		a.log.Error("Passwordni heshlashda xatolik", logger.Error(err))
		return entities.RegistrRes{}, pkgerrors.NewError(http.StatusInternalServerError, "Passwordni heshlashda xatolik")
	}
	err = a.storage.Admin().Registration(ctx, req)
	if err != nil {
		a.log.Error("Telefon raqamini saqlashda xatolik", logger.Error(err))
		return entities.RegistrRes{}, pkgerrors.NewError(http.StatusInternalServerError, "Telefon raqamini saqlashda xatolik")
	}
	tokenMetadata := map[string]string{
		"userId": req.ID,
		"role":     "user",
	}
	tokens := entities.Tokens{}
	tokens.AccessToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTAccessTokenExpireDuration, a.cfg.JWTSecretKey)
	if err != nil {
		a.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return entities.RegistrRes{}, err
	}
	tokens.RefreshToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTRefreshTokenExpireDuration, a.cfg.JWTSecretKey)
	if err != nil {
		a.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return entities.RegistrRes{}, err
	}

	a.log.Info("Registration finished")
	return entities.RegistrRes{
		ID:    req.ID,
		Token: tokens,
	}, nil
}

func (a adminController) Login(ctx context.Context, req entities.LoginReq) (entities.LoginRes, error) {
	a.log.Info("Login started: ")
     var res entities.PassId
	 var err error
	res, err = a.storage.Admin().Login(ctx, req.Username)
	if err != nil {
		a.log.Error("Passwordni olishda xatolik", logger.Error(err))
		return entities.LoginRes{}, pkgerrors.NewError(http.StatusInternalServerError, "Passwordni olishda xatolik")
	}

	isValid := utils.CheckPasswordHash(req.Password, res.Password)
	if !isValid {
		a.log.Error("Noto'g'ri parol")
		return entities.LoginRes{}, pkgerrors.NewError(http.StatusUnauthorized, "Noto'g'ri username yoki parol")
	}
	tokenMetadata := map[string]string{
		"userId": res.ID,
		"role":     "user",
	}

	tokens := entities.Tokens{}
	tokens.AccessToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTAccessTokenExpireDuration, a.cfg.JWTSecretKey)
	if err != nil {
		a.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return entities.LoginRes{}, err
	}
	tokens.RefreshToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTRefreshTokenExpireDuration, a.cfg.JWTSecretKey)
	if err != nil {
		a.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return entities.LoginRes{}, err
	}
	return entities.LoginRes{
		Username: req.Username,
		Token:    tokens,
	}, nil
}

func (a adminController) Logout(ctx context.Context, req entities.BlockedToken) error {
	a.log.Info("Logout started")

	claims, err := jwt.ParseJWT(req.Token)
	if err != nil {
		a.log.Error("Tokenni parslashda xatolik", logger.Error(err))
		return pkgerrors.NewError(http.StatusUnauthorized, "Yaroqsiz token")
	}

	expiresAt, ok := claims["exp"].(float64)
	if !ok {
		a.log.Error("Tokendagi talab qilinadigan ma'lumotlar yo'q")
		return pkgerrors.NewError(http.StatusUnauthorized, "Yaroqsiz token")
	}

	blockedToken := entities.BlockedToken{
		Token:     req.Token,
		ExpiresAt: int64(expiresAt),
	}
    
	blockedToken.ID=uuid.NewString()
	err = a.storage.Admin().Logout(ctx, blockedToken)
	if err != nil {
		a.log.Error("Tokenni bloklashda xatolik", logger.Error(err))
		return pkgerrors.NewError(http.StatusInternalServerError, "Tokenni bloklashda xatolik")
	}

	a.log.Info("Logout muvaffaqiyatli yakunlandi")
	return nil
}

func (a adminController) CreateTweet(ctx context.Context, req entities.CreateTweetReq) error{
	a.log.Info("CreateTweet started: ")

    err :=a.storage.Admin().CreateTweet(ctx, req)
	if err != nil {
		a.log.Error("Tweet yaratishda xatolik", logger.Error(err))
		return pkgerrors.NewError(http.StatusInternalServerError, "Tweet yaratishda xatolik")
	}
   a.log.Info("CreateTweet muvaffaqiyatli yakunlandi")
   return nil
}

func (a adminController) Follow(ctx context.Context, req entities.Follow) error {
	a.log.Info("Follow started: ")

	err :=a.storage.Admin().Follow(ctx, req)
	if err != nil {
		a.log.Error("Follow qilishda xatolik", logger.Error(err))
		return pkgerrors.NewError(http.StatusInternalServerError, "Follow qilishda xatolik")
	}
   a.log.Info("Follow qilish muvaffaqiyatli yakunlandi")
   return nil
}

func (a adminController) Unfollow(ctx context.Context, req entities.Follow) error {
	a.log.Info("Unfollow started: ")

	err :=a.storage.Admin().Unfollow(ctx, req)
	if err != nil {
		a.log.Error("Unfollow qilishda xatolik", logger.Error(err))
		return pkgerrors.NewError(http.StatusInternalServerError, "Unfollow qilishda xatolik")
	}
   a.log.Info("Unfollow qilish muvaffaqiyatli yakunlandi")
   return nil
}

func (a adminController) GetFollowers(ctx context.Context, userId string) ([]entities.User, error) {
	a.log.Info("GetFollowers started: ")

	followers, err :=a.storage.Admin().GetFollowers(ctx, userId)
	if err != nil {
		a.log.Error("GetFollowers xatolik", logger.Error(err))
		return []entities.User{}, pkgerrors.NewError(http.StatusInternalServerError, "")
	}
   a.log.Info("GetFollowers muvaffaqiyatli yakunlandi")
   return followers, nil
}

func (a adminController) GetFollowings(ctx context.Context, userId string) ([]entities.User, error){
	a.log.Info("GetFollowings started: ")

	followers, err :=a.storage.Admin().GetFollowings(ctx, userId)
	if err != nil {
		a.log.Error("GetFollowings qilishda xatolik", logger.Error(err))
		return []entities.User{}, pkgerrors.NewError(http.StatusInternalServerError, "")
	}
   a.log.Info("GetFollowings muvaffaqiyatli yakunlandi")
   return followers, nil
}

func (a adminController) LikeTweet(ctx context.Context, req entities.Like) error{
    a.log.Info("LikeTweet started: ")

	err :=a.storage.Admin().LikeTweet(ctx, req)
	if err != nil {
		a.log.Error("LikeTweet qilishda xatolik", logger.Error(err))
		return pkgerrors.NewError(http.StatusInternalServerError, "LikeTweet funksiyasida xatolik")
	}
   a.log.Info("LikeTweet muvaffaqiyatli yakunlandi")
   return nil
}

func (a adminController) UnlikeTweet(ctx context.Context, req entities.Like) error{
    a.log.Info("UnlikeTweet started: ")

	err :=a.storage.Admin().UnlikeTweet(ctx, req)
	if err != nil {
		a.log.Error("UnlikeTweet qilishda xatolik", logger.Error(err))
		return pkgerrors.NewError(http.StatusInternalServerError, "UnlikeTweet funksiyasida xatolik")
	}
   a.log.Info("UnlikeTweet muvaffaqiyatli yakunlandi")
   return nil
}

func (a adminController) Retweet(ctx context.Context, req entities.Retweet) error{
	a.log.Info("Retweet started: ")

	err :=a.storage.Admin().Retweet(ctx, req)
	if err != nil {
		a.log.Error("Retweet qilishda xatolik", logger.Error(err))
		return pkgerrors.NewError(http.StatusInternalServerError, "Retweet funksiyasida xatolik")
	}
   a.log.Info("Retweet muvaffaqiyatli yakunlandi")
   return nil
}

func (a adminController) Search(ctx context.Context, req string)([]entities.User, error){
	a.log.Info("Search started: ")

	users, err := a.storage.Admin().Search(ctx, req)
	if err != nil {
		a.log.Error("Retweet qilishda xatolik", logger.Error(err))
		return []entities.User{}, pkgerrors.NewError(http.StatusInternalServerError, "Retweet funksiyasida xatolik")
	}
   a.log.Info("Retweet muvaffaqiyatli yakunlandi")
   return users, nil
}