package handlers

import (
	"encoding/json"
	"fiber-ferdi/domains/model/services"
	"fiber-ferdi/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllModel(c *fiber.Ctx) error {
    search := c.Query("search_term")
	currentPage, _ := strconv.Atoi(c.Query("page_number"))
	if currentPage <= 0 {
		currentPage = 1
	}
	itemsPerPage, _ := strconv.Atoi(c.Query("page_size"))
	if itemsPerPage <= 0 {
		itemsPerPage = 10
	}


    code, msg, data , totalItems := services.GetAllModel(search, currentPage, itemsPerPage)
	totalPages := totalItems / itemsPerPage
	if totalItems % itemsPerPage != 0 {
		totalPages++ // Tambahkan satu halaman jika ada sisa item
	}
		paginationHeader := fiber.Map{
			"total_count":    strconv.Itoa(totalItems),
			"page_size": strconv.Itoa(itemsPerPage),
			"page_number":   strconv.Itoa(currentPage),
			"total_pages":   strconv.Itoa(totalPages),
		}

		jsonHeader, err := json.Marshal(paginationHeader)
		if err != nil {
			// log.Fatal(err)
			panic(err)
		}
		c.Set("X-Pagination", string(jsonHeader))
		return c.Status(code).JSON(fiber.Map{
			"code":    code,
			"message": msg,
			"data":    data,
		})
}


func GetModelById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	code, msg, data := services.GetModelById(id)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

func StoreModel(c *fiber.Ctx) error {

	json := new(models.Model)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid JSON",
		})
	}
	code, msg := services.StoreModel(c.Locals("user_id").(int) ,json )
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func UpdateModel(c *fiber.Ctx) error {
	json := new(models.Model)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid JSON",
		})
	}
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	code, msg := services.UpdateModel(id, json)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func DeleteModel(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	code, msg := services.DeleteModel(id)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}
