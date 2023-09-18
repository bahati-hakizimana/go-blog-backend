package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bahati-hakizimana/blogue-backend/controller"
	 "github.com/bahati-hakizimana/blogue-backend/middleware"

)


func Setup(app *fiber.App){
	app.Post("/api/register",controller.Register)
	app.Post("/api/login",controller.Login)
     app.Use(middleware.IsAuthanticate)
	 app.Post("/api/post",controller.CreatePost)
	 app.Get("/api/allpost",controller.AllPost)
	 app.Get("/api/allpost/:id",controller.DetailPost)
	 app.Put("/api/updatepost/:id",controller.UpdatePost)
	 app.Get("/api/uniquepost",controller.UniquePost)
	 app.Delete("/api/deletepost/:id",controller.DeletePost)
	 app.Post("/api/upload-image",controller.Upload)
}