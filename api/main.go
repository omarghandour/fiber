package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}
func handler() http.HandlerFunc {
	app := fiber.New()

	app.Get("/v1", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"version": "v1",
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello friend")
	})
	app.Get("/v2", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"version": "v2",
		})
	})

	// app.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.JSON(fiber.Map{
	// 		"uri":  ctx.Request().URI().String(),
	// 		"path": ctx.Path(),
	// 	})
	// })

	return adaptor.FiberApp(app)
}
