package repositories

import (
	"fiber-ferdi/configs"
	"fiber-ferdi/models"
	"fiber-ferdi/utils"
	"fmt"
)

func CheckUser(username string) (*models.User, error) {
	user := models.User{}
	query := models.User{Username: username}
	db := configs.DB
	result := db.Preload("Role").First(&user, &query)
	// fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func Register(json *models.User) error {
	fmt.Println(json)
	user := models.User{Username: json.Username, Password: utils.HashAndSalt([]byte(json.Password)), RoleID: json.RoleID}
	db := configs.DB
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func Register(json *models.User) error {
// 	user := models.User{
//     Username: json.Username,
//     Password: utils.HashAndSalt([]byte(json.Password)),
//     // Role:   json.Role,
// }
