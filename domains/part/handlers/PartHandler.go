package handlers

import (
	"encoding/json"
	"fiber-ferdi/domains/part/services"
	"fiber-ferdi/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllPart(c *fiber.Ctx) error {
    search := c.Query("search_term")
	currentPage, _ := strconv.Atoi(c.Query("page_number"))
	if currentPage <= 0 {
		currentPage = 1
	}
	itemsPerPage, _ := strconv.Atoi(c.Query("page_size"))
	if itemsPerPage <= 0 {
		itemsPerPage = 10
	}


    code, msg, data , totalItems := services.GetAllPart(search, currentPage, itemsPerPage)
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
		// log.Println(data[len(data)-1])
		c.Set("X-Pagination", string(jsonHeader))
		return c.Status(code).JSON(fiber.Map{
			"code":    code,
			"message": msg,
			"data":    data,
		})
}
func ExportExcelHandler(c *fiber.Ctx) error {
    // Panggil fungsi GetExcle untuk menghasilkan file Excel
    data := services.GetExcle(c)
    if data != nil {
        // Jika terjadi error, tangani di sini
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code": 500,
            "msg":  "Internal Server Error",
            "data": nil,
        })
    }
return data
    // Jika tidak ada error, Anda bisa mengirim respons sukses
    // return c.JSON(fiber.Map{
    //     "code": 200,
    //     "msg":  "Success",
    //     "data":  data , // Tidak ada data yang perlu dikirim karena file Excel dikirim sebagai respons
    // })
}


func GetPartById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	code, msg, data := services.GetPartById(id)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

func StorePart(c *fiber.Ctx) error {

	json := new(models.Part)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid JSON",
		})
	}
	code, msg := services.StorePart(c.Locals("user_id").(int) ,json )
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func UpdatePart(c *fiber.Ctx) error {
	json := new(models.Part)
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
	code, msg := services.UpdatePart(id, json)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func DeletePart(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	code, msg := services.DeletePart(id)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}
