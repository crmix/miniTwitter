package routers

import (
	"miniTwitter/configs"
	"miniTwitter/handlers"
	"miniTwitter/logger"
	"miniTwitter/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler     handlers.Handler
	config      *configs.Configuration
	router      *gin.Engine
	logger      logger.LoggerI
	middlewares *middlewares.JWTRoleAuthorizer
}

// New creates a new router
func New(h handlers.Handler, cfg *configs.Configuration, logger logger.LoggerI, mw *middlewares.JWTRoleAuthorizer) Router {
	r := gin.New()

	return Router{
		handler:     h,
		router:      r,
		logger:      logger,
		config:      cfg,
		middlewares: mw,
	}

}

func (r Router) Start() {

	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())
//	r.router.Use(middlewares.CustomCORSMiddleware())

	r.UserRouters()

	r.logger.Info("HTTP: Server being started...", logger.String("port", r.config.HTTPPort))

	err := r.router.Run(r.config.HTTPPort)
	if err != nil {
		panic(err)
	}
}
