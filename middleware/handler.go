// package error

// import(
// 	"net/http"
// 	"github.com/gofiber/fiber/v2"

// )
//
//	func ErrorHandling(c *fiber.Ctx, err error)error{
//		return c.Status(status : http.StatusInternalServerError).SendString(body :"internal server Error")
//	}
package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Recovery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Println("datatat")
				c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
			}
		}()
		return c.Next()
	}
}
