package routers

func(r Router) UserRouters() {
	apiGroup := r.router.Group("/api")
	apiGroup.POST("/auth/registr", r.handler.Registration)
	apiGroup.POST("/auth/login", r.handler.Login)
	apiGroup.POST("/auth/logout", r.handler.Logout)
  
	apiGroup.POST("/tweets/", r.handler.CreateTweet)
	apiGroup.POST("/follow", r.handler.Follow)
	apiGroup.POST("/unfollow", r.handler.Unfollow)
	apiGroup.GET("/followers", r.handler.GetFollowers)
	apiGroup.GET("/followings", r.handler.GetFollowings)

	apiGroup.POST("/like", r.handler.LikeTweet)
	apiGroup.POST("/unlike", r.handler.UnlikeTweet)
	apiGroup.POST("/retweet", r.handler.Retweet)
	apiGroup.GET("/search", r.handler.Search)
}