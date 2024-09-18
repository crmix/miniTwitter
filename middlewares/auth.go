package middlewares

import (
	"errors"
	"fmt"
	"log"
	"miniTwitter/configs"
	"miniTwitter/logger"
	"miniTwitter/pkg/jwt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JWTRoleAuthorizer struct {
	enforcer interface {
		Enforce(rvals ...interface{}) (bool, error)
	}
	signingKey []byte
	logger     logger.LoggerI
}

// NewCasbinJWTRoleAuthorizer creates and returns a new Role Authorizer
func NewCasbinJWTRoleAuthorizer(cfg *configs.Configuration, logger logger.LoggerI) (*JWTRoleAuthorizer, error) {
	enforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, cfg.MiddlewareRolesPath)
	if err != nil {
		logger.Fatal("could not initialize new enforcer", zap.Any("error", err))
		return nil, err
	}

	return &JWTRoleAuthorizer{
		enforcer:   enforcer,
		signingKey: []byte(cfg.JWTSecretKey),
		logger:     logger,
	}, nil
}

// Middleware checks the permission using Casbin
func (jwta *JWTRoleAuthorizer) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowed, err := jwta.checkPermission(c.Request)
		if err != nil {
			jwta.logger.Error("Error checking permission", logger.Error(err))
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !allowed {
			jwta.logger.Info("Error checking permission: not allowed")
			c.AbortWithError(http.StatusForbidden, errors.New("permission denied"))
			return
		}
		c.Next()
	}
}

func (jwta *JWTRoleAuthorizer) checkPermission(r *http.Request) (bool, error) {
	role, err := jwta.getRole(r.Header.Get("Authorization"))
	if err != nil {
		return false, err
	}

	method := r.Method
	path := r.URL.Path
	enforced, err := jwta.enforcer.Enforce(role, path, method)
	return enforced, err
}

func (jwta *JWTRoleAuthorizer) getRole(accessToken string) (string, error) {
	role, err := jwt.ExtractFromClaims("role", accessToken, jwta.signingKey)
	if err != nil {
		log.Println("could not extract claims:", err)
		return "", err
	}

	if _, ok := role.(string); !ok {
		return "", fmt.Errorf("role: %v not stringable", role)
	}

	return role.(string), nil
}
