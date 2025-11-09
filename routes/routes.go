package routes

import (
	"github.com/gofiber/fiber/v2"
	"news-feed-system-backend/app/handlers"
	"news-feed-system-backend/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	public := api.Group("")
	public.Post("/register", handlers.Register)
	public.Post("/login", handlers.Login)

	private := api.Group("", middleware.AuthRequired)
	private.Post("/posts", handlers.CreatePost)
	private.Get("/feed", handlers.GetFeed)
	private.Get("/users", handlers.GetAllUsers)
	private.Get("/users/:userid", handlers.GetUserByID)
	private.Post("/follow/:userid", handlers.Follow)
	private.Delete("/follow/:userid", handlers.Unfollow)
	private.Get("/following", handlers.GetFollowing)
}
