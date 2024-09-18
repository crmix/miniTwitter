package handlers

import (
	"net/http"

	"miniTwitter/configs"
	adminController "miniTwitter/controllers"
	"miniTwitter/logger"
	e "miniTwitter/pkg/errors"

	"github.com/gin-gonic/gin"

	httppkg "miniTwitter/pkg/http"
)

type Handler struct {
	cfg             *configs.Configuration
	log             logger.LoggerI
	adminController adminController.AdminController
}

func New(
	cfg *configs.Configuration,
	log logger.LoggerI,
	adminController adminController.AdminController,
) Handler {
	return Handler{
		cfg:             cfg,
		log:             log,
		adminController: adminController,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status httppkg.Status, data ...interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, httppkg.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

// StatusFromError ...
func StatusFromError(err error) httppkg.Status {
	if err == nil {
		return httppkg.OK
	}

	code, ok := e.ExtractStatusCode(err)
	if !ok || code == http.StatusInternalServerError {
		return httppkg.Status{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: err.Error(),
		}
	} else if code == http.StatusNotFound {
		return httppkg.Status{
			Code:        http.StatusNotFound,
			Status:      "NOT_FOUND",
			Description: err.Error(),
		}
	} else if code == http.StatusBadRequest {
		return httppkg.Status{
			Code:        http.StatusBadRequest,
			Status:      "BAD_REQUEST",
			Description: err.Error(),
		}
	} else if code == http.StatusForbidden {
		return httppkg.Status{
			Code:        http.StatusForbidden,
			Status:      "FORBIDDEN",
			Description: err.Error(),
		}
	} else if code == http.StatusUnauthorized {
		return httppkg.Status{
			Code:        http.StatusUnauthorized,
			Status:      "FORBIDDEN",
			Description: err.Error(),
		}
	} else {
		return httppkg.Status{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: err.Error(),
		}
	}

}