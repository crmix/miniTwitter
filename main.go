package main

import (
	"miniTwitter/configs"
	"miniTwitter/constants"
	"miniTwitter/controllers"
	"miniTwitter/handlers"
	"miniTwitter/logger"
	"miniTwitter/middlewares"
	"miniTwitter/routers"
	"miniTwitter/storage"

	"github.com/gin-gonic/gin"
)

func main(){
	cfg := configs.Config()

	strg := storage.New(cfg)

	switch cfg.Environment {
	case constants.DebugMode:
		gin.SetMode(gin.DebugMode)
	case constants.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	//logger
	log := logger.NewLogger(cfg.AppName, cfg.LogLevel)
	defer logger.Cleanup(log)


	admincontroller := controllers.NewAdminController(log, strg)

	h := handlers.New(
		cfg,
		log,
		admincontroller,
	)
	router := routers.New(h, cfg, log, &middlewares.JWTRoleAuthorizer{})

	router.Start()
}