package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
	Role     string
}

func main() {
	app := fiber.New()
	app.Get("/post", handlerGetPost)
	app.Get("/post/manage", onlyAdmin(handlerGetPostManage))
	app.Get("/post/special", onlySpecialUser(handlerGetPostSpecial))
	log.Fatal(app.Listen(":4000"))
}

func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := getUserFromDB()

		if user.Role != "admin" {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return fn(c)
	}
}

func onlySpecialUser(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := getUserFromDB()
		if user.Role == "special" || user.Role == "admin" {
			return fn(c)
		}
		return fn(c)
	}
}

func getUserFromDB() User {
	return User{
		Username: "James",
		Role:     "special",
	}

}

func handlerGetPost(c *fiber.Ctx) error {
	return c.JSON("some posts here")
}

func handlerGetPostManage(c *fiber.Ctx) error {
	return c.JSON("the admin page of this post")
}
func handlerGetPostSpecial(c *fiber.Ctx) error {
	return c.JSON("the special page of this post")
}
