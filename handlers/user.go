package handlers

import (
	"miniTwitter/constants"
	"miniTwitter/entities"
	"miniTwitter/logger"
	htp "miniTwitter/pkg/http"
	jwta "miniTwitter/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Registration(c *gin.Context) {
	var req entities.RegistrReq
	err := c.ShouldBindJSON(&req)

	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}
	req.ID = uuid.NewString()

	resp, err := h.adminController.Registration(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, resp)
}
func (h *Handler) Login(c *gin.Context) {
	var req entities.LoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}

	resp, err := h.adminController.Login(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, resp)
}
func (h *Handler) Logout(c *gin.Context) {
	var req entities.BlockedToken
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}
	err = h.adminController.Logout(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, constants.SuccessMessage)
}
func (h *Handler) CreateTweet(c *gin.Context) {
	userId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
		return
	}

	var req entities.CreateTweetReq
	err = c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}
	req.UserId = userId
	req.ID = uuid.NewString()
	err = h.adminController.CreateTweet(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, constants.SuccessMessage)
}

// func (h *Handler) GetTweets(c *gin.Context) {
// 	userId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
// 	if err != nil {
// 		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
// 		return
// 	}
// 	tweets, err := h.adminController.GetTweets(c, userId)
// 	if err !=nil {
// 		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
// 		return
// 	}
// 	h.handleResponse(c, htp.OK, tweets)

// }

func (h *Handler) Follow(c *gin.Context) {
	followerId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
		return
	}
	var req entities.Follow
	err = c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}
	req.FollowerID = followerId
	req.ID = uuid.NewString()
	err = h.adminController.Follow(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, constants.SuccessMessage)
}

func (h *Handler) Unfollow(c *gin.Context) {
	followerId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
		return
	}
	var req entities.Follow
	err = c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}
	req.FollowerID = followerId
	err = h.adminController.Unfollow(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, constants.SuccessMessage)
}

func (h *Handler) GetFollowers(c *gin.Context) {
	userId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
		return
	}
	followers, err := h.adminController.GetFollowers(c, userId)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, followers)
}

func (h *Handler) GetFollowings(c *gin.Context) {
	userId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
		return
	}

	followings, err := h.adminController.GetFollowings(c, userId)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, followings)
}

func (h *Handler) LikeTweet(c *gin.Context) {
	var req entities.Like

	userId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}
	req.ID = uuid.NewString()
	req.UserId = userId
	err = h.adminController.LikeTweet(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, constants.SuccessMessage)
}

func (h *Handler) UnlikeTweet(c *gin.Context) {
	var req entities.Like

	userId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}
	req.ID = uuid.NewString()
	req.UserId = userId
	err = h.adminController.UnlikeTweet(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, constants.SuccessMessage)
}

func (h *Handler) Retweet(c *gin.Context) {
	var req entities.Retweet
	userId, err := jwta.ExtractUserIDFromToken(c, []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, "Invalid or expired token")
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, logger.Error(err))
		return
	}
	req.UserID = userId
	req.OriginalTweetID = req.ID
	req.ID = uuid.NewString()
	err = h.adminController.Retweet(c, req)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, constants.SuccessMessage)

}

func (h *Handler) Search(c *gin.Context) {
	searchQuery := c.Query("q")
	if searchQuery == "" {
		h.handleResponse(c, htp.BadRequest, "Search query cannot be ampty")
		return
	}
	data, err := h.adminController.Search(c, searchQuery)
	if err != nil {
		h.handleResponse(c, htp.InternalServerError, logger.Error(err))
		return
	}
	h.handleResponse(c, htp.OK, data)
}
