package constants

import "time"

const (
	TestMode  = "test"
	DebugMode = "debug"

	JWTRefreshTokenExpireDuration = time.Hour * 72
	JWTAccessTokenExpireDuration  = time.Minute * 60
	ContextTimeoutDuration        = time.Second * 7

	SuccessMessage = "Success"
)