package main

import (
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  app.Get("/projects", func(c *fiber.Ctx) error {
    return c.JSON([]map[string]string{
      {"title": "Portfolio", "tech": "Go + React"},
      {"title": "Blog", "tech": "Next.js"},
    })
  })

  app.Listen(":8000")
}
