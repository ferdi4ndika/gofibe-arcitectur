package services

import (
	"fiber-ferdi/domains/auth/repositories"
	"fiber-ferdi/models"
	"fiber-ferdi/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(json *models.User) (int, string, string) {
	check, err := repositories.CheckUser(json.Username)
	if err != nil {
		return fiber.StatusBadRequest, "Login failed, because username has not been registered", ""
	}
	result := utils.ComparePass(check.Password, []byte(json.Password))
	if result {
		token := utils.GenerateToken(json.Username)
		return fiber.StatusOK, "Login success", token
	}
	return 400 , "Login failed", ""
}

// func Register(json *models.User) (int, string) {
		
// 	_, err_ := repositories.CheckUser(json.Username)
// 	if err_ == nil {
// 		return fiber.StatusBadRequest, "Register failed, because username is already exist!"
// 	}
// 	err := repositories.Register(json)
// 	if err != nil {
// 		log.Fatal(err)
// 		return fiber.StatusBadRequest, "Register failed"
// 	}
// 	return fiber.StatusOK, "Register success"
// }
func Register(json *models.User) (int, string) {
    _, errCheckUser := repositories.CheckUser(json.Username)
    if errCheckUser == nil {
        return fiber.StatusBadRequest, "Register failed, because username already exists!"
    }

    err := repositories.Register(json)
    if err != nil {
        errorMessage := fmt.Sprintf("Error registering user: %v", err)
        // Di sini, Anda dapat melakukan tindakan yang sesuai, misalnya mengembalikan pesan kesalahan
        return fiber.StatusInternalServerError, errorMessage
    }

    return fiber.StatusOK, "Register success"
}

