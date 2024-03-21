package middleware

import (
	"fiber-ferdi/configs"
	"fiber-ferdi/models"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Middleware Authentication dengan izin yang diharapkan sebagai parameter
func Authentication(permission string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        authorizationHeader := c.Request().Header.Peek("Authorization")
        if !strings.Contains(string(authorizationHeader), "Bearer") {
            return c.JSON(fiber.Map{
                "code":    401,
                "message": "Invalid Token",
            })
        }

        tokenString := strings.Replace(string(authorizationHeader), "Bearer ", "", -1)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("AC9D72E4B61B7935B28D412F72193"), nil
        })
        if err != nil || !token.Valid {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "code":    401,
                "message": "Invalid Token",
            })
        }

        claims := token.Claims.(jwt.MapClaims)
        user := new(models.User)
        db := configs.DB
        query := models.User{Username: claims["username"].(string)}
        if err := db.Preload("Role").First(user, query).Error; err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "code":    fiber.StatusInternalServerError,
                "message": "Failed to load user role",
            })
        }

        permissions := strings.Split(user.Role.Permessions, ", ")
        fmt.Println(permissions)
      
		 

        // Periksa izin yang diberikan dalam parameter
        hasPermission := false
        for _, p := range permissions {
            if p == permission {
                hasPermission = true
                break
            }
        }

        if !hasPermission {
            return c.JSON(fiber.Map{
                "code":    401,
                "message": "Not Access",
            })
        }
        c.Locals("user_id", user.ID)
        return c.Next()
    }
}

// func Authentication(c *fiber.Ctx) error {
// 	authorizationHeader := c.Request().Header.Peek("Authorization")
// 	if !strings.Contains(string(authorizationHeader), "Bearer") {
// 		return c.JSON(fiber.Map{
// 			"code":    401,
// 			"message": "Invalid Token",
// 		})
// 	}
// 	tokenString := strings.Replace(string(authorizationHeader), "Bearer ", "", -1)
//     token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("AC9D72E4B61B7935B28D412F72193"), nil
// 	})
// 	if err != nil || !token.Valid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"code":    401,
// 			"message": "Invalid Token",
// 		})
// 	}
// 	claims := token.Claims.(jwt.MapClaims)
// 	user := new(models.User)
//     db := configs.DB
//     query := models.User{Username: claims["username"].(string)}
//     if err := db.Preload("Role").First(user, query).Error; err != nil {
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "code":    fiber.StatusInternalServerError,
//             "message": "Failed to load user role",
//         })
//     }
//       permissions := strings.Split(user.Role.Permessions, ", ")

// 	ok := utils.VerifyToken(tokenString)
// 	if !ok {
// 		// return c.Next()
// 		return c.JSON(fiber.Map{
// 		"code":    401,
// 		"message": "Invalid Token",
// 	})
// 	}
	
// 	hasDataPermission := false
// 	for _, permission := range permissions {
// 		if permission == "Dataaa" {
// 			hasDataPermission = true
// 			break
// 		}
// 	}

// 	// Jika izin pengguna termasuk "Data", lanjutkan ke handler berikutnya
// 	if hasDataPermission {
// 		return c.Next()
// 	}

// 	// Jika tidak, kembalikan respons JSON bahwa pengguna tidak diizinkan
// 	return c.JSON(fiber.Map{
// 		"code":    401,
// 		"message": "Not Access",
// 	})
// }
