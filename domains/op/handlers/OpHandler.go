package handlers

import (
	"encoding/json"
	"fiber-ferdi/domains/op/services"
	"fiber-ferdi/models"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllOp(c *fiber.Ctx) error {
	defer func() {
            if r := recover(); r != nil {
                c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "error": "Internal Server Error",
                })
            }
        }()
    search := c.Query("search_term")
	currentPage, _ := strconv.Atoi(c.Query("page_number"))
	if currentPage <= 0 {
		currentPage = 1
	}
	itemsPerPage, _ := strconv.Atoi(c.Query("page_size"))
	if itemsPerPage <= 0 {
		itemsPerPage = 10
	}
    code, msg, data , totalItems := services.GetAllOp(search, currentPage, itemsPerPage)
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
			panic(err)
		}
		c.Set("X-Pagination", string(jsonHeader))
		return c.Status(code).JSON(fiber.Map{
			"code":    code,
			"message": msg,
			"data":    data,
		})
}


func GetOpById(c *fiber.Ctx) error {
    defer func() {
        if r := recover(); r != nil {
            c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Internal Server Error",
            })
			}
		}()
		param := c.Params("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			panic(err) // Mengubah log.Fatal menjadi panic
		}
		code, msg, data := services.GetOpById(id)
		return c.Status(code).JSON(fiber.Map{
			"code":    code,
			"message": msg,
			"data":    data,
		})
}


func StoreOp(c *fiber.Ctx) error {
	defer func() {
				if r := recover(); r != nil {
					c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Internal Server Error",
					})
				}
			}()
		json := new(models.Op)
		if err := c.BodyParser(json); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": "Invalid JSON",
			})
		}
		code, msg := services.StoreOp(c.Locals("user_id").(int) ,json )
		return c.Status(code).JSON(fiber.Map{
			"code":    code,
			"message": msg,
		})
}

func UpdateOp(c *fiber.Ctx) error {
	defer func() {
            if r := recover(); r != nil {
                c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "error": "Internal Server Error",
                })
            }
        }()
		json := new(models.Op)
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
		code, msg := services.UpdateOp(id, json)
		return c.Status(code).JSON(fiber.Map{
			"code":    code,
			"message": msg,
		})
}

func DeleteOp(c *fiber.Ctx) error {
		defer func() {
            if r := recover(); r != nil {
                c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                    "error": "Internal Server Error",
                })
            }
        }()
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	code, msg := services.DeleteOp(id)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}
