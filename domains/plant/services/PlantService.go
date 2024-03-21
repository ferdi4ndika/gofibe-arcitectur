package services

import (
	"fiber-ferdi/domains/plant/repositories"
	"fiber-ferdi/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// func GetAllPlant(search string) (int, string, []models.Plant) {
// 	result, err := repositories.GetAllPlant(search )
// 	if err != nil {
// 		log.Fatal(err)
// 		return fiber.StatusBadRequest, err.Error(), nil
// 	}
// 	return fiber.StatusOK, "Get data plant success", result
// }
func GetAllPlant(search string, currentPage, itemsPerPage int) (int, string, []models.Plant, int) {
    // Mendapatkan data plant dari repository dengan menggunakan search, currentPage, dan itemsPerPage
    result, totalCount, err := repositories.GetAllPlant(search, currentPage, itemsPerPage)
    if err != nil {
        log.Fatal(err)
        return fiber.StatusBadRequest, err.Error(), nil, 0
    }
    return fiber.StatusOK, "Get data plant success", result, totalCount
}

func GetAllPlantWs() ([]map[string]interface{}, error) {
    plants, err := repositories.GetAllPlantWs()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    var plantsArray []map[string]interface{}

    for _, plant := range plants {
        // total := plant.Qty * plant.Price
        plantData := map[string]interface{}{
            "id":  plant.ID,
            "code":  plant.Code,
            "information":   plant.Information,
            "created_at": plant.CreatedAt,
            "updated_at":plant.CreatedAt,
        }
        plantsArray = append(plantsArray, plantData)
    }

    return plantsArray, nil
}
func GetPlantById(id int) (int, string, models.Plant) {
	result, err := repositories.GetPlantById(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error(), models.Plant{}
	}
	return fiber.StatusOK, "Get data plant success", result
}

func StorePlant(user_id int ,json *models.Plant) (int, string) {
	err := repositories.StorePlant(user_id ,json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Insert data plant success"
}

func UpdatePlant(id int, json *models.Plant) (int, string) {
	err := repositories.UpdatePlant(id, json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Update data plant success"
}

func DeletePlant(id int) (int, string) {
	err := repositories.DeletePlant(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Delete data plant success"
}
