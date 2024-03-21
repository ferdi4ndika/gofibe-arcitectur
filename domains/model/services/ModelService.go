package services

import (
	"fiber-ferdi/domains/model/repositories"
	"fiber-ferdi/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllModel(search string, currentPage, itemsPerPage int) (int, string, []models.Model, int) {
    // Mendapatkan data Model dari repository dengan menggunakan search, currentPage, dan itemsPerPage
    result, totalCount, err := repositories.GetAllModel(search, currentPage, itemsPerPage)
    if err != nil {
        log.Fatal(err)
        return fiber.StatusBadRequest, err.Error(), nil, 0
    }
    return fiber.StatusOK, "Get data Model success", result, totalCount
}

func GetModelById(id int) (int, string, models.Model) {
	result, err := repositories.GetModelById(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error(), models.Model{}
	}
	return fiber.StatusOK, "Get data Model success", result
}

func StoreModel(user_id int ,json *models.Model) (int, string) {
	err := repositories.StoreModel(user_id ,json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Insert data Model success"
}

func UpdateModel(id int, json *models.Model) (int, string) {
	err := repositories.UpdateModel(id, json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Update data Model success"
}

func DeleteModel(id int) (int, string) {
	err := repositories.DeleteModel(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Delete data Model success"
}
