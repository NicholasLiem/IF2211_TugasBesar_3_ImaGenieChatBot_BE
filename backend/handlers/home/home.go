package home
import "github.com/gofiber/fiber/v2"

func GetHome(c *fiber.Ctx) error{
	return c.SendFile("../../../frontend/src/App.js")
}