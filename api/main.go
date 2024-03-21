package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}

func main() {
	// fmt.Println("Hello World!")
	app := fiber.New()
	todos := []Todo{}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello friend")
	})
    app.Post("/", func(c *fiber.Ctx) error {
		todo := &Todo{}

		
		if err := c.BodyParser(todo); err != nil {
			
			return err
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.JSON(todos)
	})
	log.Fatal(app.Listen(":3000"))
}
