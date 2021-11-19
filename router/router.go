package router

import (
	."github.com/erdoganakdeniz/handlers"
	."github.com/erdoganakdeniz/config"
	."github.com/erdoganakdeniz/middlewares"
	"github.com/gofiber/fiber"
)

func SetupRouter(app *fiber.App) {
	router:=app.Group("/api/v1")
	_authHandler:=AuthHandler{UsersColl:Mongo.DB.Collection("users") }
	//Auth Routes
	router.Post("/signup",_authHandler.Signup)
	router.Post("/login",_authHandler.Login)
	//User Routes
	_userHandler:=UserHandler{UserColl: Mongo.DB.Collection("users")}
	router.Get("/user",WithGuard,WithUser,_userHandler.GetUser)
	router.Put("/user",WithGuard,WithUser,_userHandler.UpdateUser)
	router.Post("/user/:id",WithGuard,WithUser,_userHandler.FollowUnFollowUser)

	//Post Routes
	_postHandler := PostHandler{
		UserColl:    Mongo.DB.Collection("users"),
		PostColl:    Mongo.DB.Collection("posts"),
		CommentColl: Mongo.DB.Collection("comments"),
	}
	router.Post("/post", WithGuard, WithUser, _postHandler.CreatePost)
	router.Put("/post/:id", WithGuard, WithUser, _postHandler.UpdatePost)
	router.Delete("/post/:id", WithGuard, WithUser, _postHandler.DeletePost)
	router.Post("/post/:id", WithGuard, WithUser, _postHandler.LikeDislikePost)
	router.Get("/post/timeline/user/:userId", _postHandler.UserTimeline)
	router.Get("/post/timeline/home/:userId?", _postHandler.HomeTimeline)

	// Comment Routes
	_commentHandler := CommentHandler{
		CommentColl: Mongo.DB.Collection("comments"),
		PostColl:    Mongo.DB.Collection("posts"),
	}
	router.Get("/comment", _commentHandler.GetComment)
	router.Post("/comment", WithGuard, WithUser, _commentHandler.CommentPost)
	router.Put("/comment/:id", WithGuard, WithUser, _commentHandler.UpdateComment)
	router.Delete("/comment/:id", WithGuard, WithUser, _commentHandler.DeleteComment)
	router.Post("/comment/:id", WithGuard, WithUser, _commentHandler.LikeDislikeComment)
}