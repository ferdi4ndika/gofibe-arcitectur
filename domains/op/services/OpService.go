package services

import (
	"fiber-ferdi/domains/op/repositories"
	"fiber-ferdi/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllOp(search string, currentPage, itemsPerPage int) (int, string, []models.Op, int) {
    // Mendapatkan data Op dari repository dengan menggunakan search, currentPage, dan itemsPerPage
    result, totalCount, err := repositories.GetAllOp(search, currentPage, itemsPerPage)
    if err != nil {
        log.Fatal(err)
        return fiber.StatusBadRequest, err.Error(), nil, 0
    }
    return fiber.StatusOK, "Get data Op success", result, totalCount
}

func GetOpById(id int) (int, string, models.Op) {
	result, err := repositories.GetOpById(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error(), models.Op{}
	}
	return fiber.StatusOK, "Get data Op success", result
}

func StoreOp(user_id int ,json *models.Op) (int, string) {
	err := repositories.StoreOp(user_id ,json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Insert data Op success"
}

func UpdateOp(id int, json *models.Op) (int, string) {
	err := repositories.UpdateOp(id, json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Update data Op success"
}

func DeleteOp(id int) (int, string) {
	err := repositories.DeleteOp(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Delete data Op success"
}
