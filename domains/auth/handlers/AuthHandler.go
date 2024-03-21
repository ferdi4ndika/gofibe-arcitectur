package handlers

import (
	"fiber-ferdi/configs"
	"fiber-ferdi/domains/auth/services"
	"fiber-ferdi/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// func Login(c *fiber.Ctx) error {
// 	var token string
// 	json := new(models.User)
// 	if err := c.BodyParser(json); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"code":    fiber.StatusBadRequest,
// 			"message": "Invalid JSON",
// 		})
// 	}

// 	code, msg, token := services.Login(json)
// 	user := new(models.User)
// 	db := configs.DB
// 	query := models.User{Username: json.Username}
//     if err := db.Preload("Role").First(user, query).Error; err != nil {
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "code":    fiber.StatusInternalServerError,
//             "message": "Failed to load user role",
//         })
//     }

// 	c.Locals("permesion", strings.Split(user.Role.Permessions, ", "))
// 	return c.Status(code).JSON(fiber.Map{
// 		"code":    code,
// 		"message": msg,
// 		"token": token,
// 		"id":    user.ID,
// 		"role":    user.Role.Role,
// 		"permesion": strings.Split(user.Role.Permessions, ", ")   ,
// 		// "permesion":    json.Role.Permessions,

// 		// "Roel": token,
// 		// "data": M{
// 		// 	"token": token,
// 		// },
// 	})}
func Login(c *fiber.Ctx) error {
    // Parsing data pengguna dari JSON
    json := new(models.User)
    if err := c.BodyParser(json); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": "Invalid JSON",
        })
    }

    // Lakukan login dan peroleh token
    code, msg, token := services.Login(json)
    
    // Ambil pengguna dari basis data
    user := new(models.User)
    db := configs.DB
    query := models.User{Username: json.Username}
    if err := db.Preload("Role").First(user, query).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "Failed to load user role",
        })
    }
    
    // Simpan izin pengguna dalam c.Locals

    
    // Kirim respons JSON yang mencakup token dan informasi pengguna
    return c.Status(code).JSON(fiber.Map{
        "code":       code,
        "message":    msg,
        "token":      token,
        "id":         user.ID,
        "role":       user.Role.Role,
        "permissions": strings.Split(user.Role.Permessions, ", "),
    })
}

func Register(c *fiber.Ctx) error {

	json := new(models.User)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid JSON",
		})
	}
	code, msg := services.Register(json)
	return c.JSON(fiber.Map{
		"code":    code,
		"message": msg,
		"data":    json,
	})
}

type M map[string]interface{}
