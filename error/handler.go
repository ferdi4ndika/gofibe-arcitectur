// package error

// import(
// 	"net/http"
// 	"github.com/gofiber/fiber/v2"

// )
//
//	func ErrorHandling(c *fiber.Ctx, err error)error{
//		return c.Status(status : http.StatusInternalServerError).SendString(body :"internal server Error")
//	}
package error

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandling(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
}
