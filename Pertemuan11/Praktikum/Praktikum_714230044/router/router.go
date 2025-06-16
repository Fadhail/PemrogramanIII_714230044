package router

import (
	"itsbackend/config/middleware"
	"itsbackend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", handler.Homepage)
	api.Get("/mahasiswa", handler.GetAllMahasiswa)
	api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM)
	api.Post("/mahasiswa", middleware.Middlewares("admin"), handler.InsertMahasiswa)
	api.Put("/mahasiswa/:npm", middleware.Middlewares("admin"), handler.UpdateMahasiswa)
	api.Delete("/mahasiswa/delete/:npm", middleware.Middlewares("admin"), handler.DeleteMahasiswa)
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}
