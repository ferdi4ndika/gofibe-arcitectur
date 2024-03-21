package handlers

import (
	"encoding/json"
	"fiber-ferdi/domains/plant/services"
	"fiber-ferdi/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)
func HandleWebSocket(c *fiber.Ctx) error {
    // membuak websoket
	if err := websocket.New(func(conn *websocket.Conn) {
    defer conn.Close()

    // run real time
    for {
        // get data
        plants, err := services.GetAllPlantWs()
        if err != nil {
            log.Println("Error getting data from service:", err)
            continue
        }

        // membaca konesi
        if err := conn.WriteJSON(plants); err != nil {
            log.Println("Error writing JSON to WebSocket:", err)
            continue
        }

        // waktu untuk query
        time.Sleep(1 * time.Second)
    }
})(c); err != nil {
    log.Println("Web shoket tidak update:", err)
    return err
}
    return nil
}

// func GetAllPlant(c *fiber.Ctx) error {
// 	search := c.Query("search_term")
// 	code, msg, data := services.GetAllPlant(search)
// 	return c.Status(code).JSON(fiber.Map{
// 		"code":    code,
// 		"message": msg,
// 		"data":    data,
// 	})
// }

func GetAllPlant(c *fiber.Ctx) error {
	defer func() {
			if r := recover(); r != nil {
				// Tangani error di sini
				c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
			}
		}()
		// panic("dasasa")
    search := c.Query("search_term")
    // itemsPerPage := 1  
    // currentPage := 1   
	// currentPageStr := c.Query("page_number")
    // currentPage, err := strconv.Atoi(currentPageStr)
    // if err != nil || currentPage <= 0 {
    //     currentPage = 1
    // }
    // itemsPerPageStr := c.Query("page_size")
    // itemsPerPage, err := strconv.Atoi(itemsPerPageStr)
    // if err != nil || itemsPerPage <= 0 {
    //     itemsPerPage = 10
    // }
	currentPage, _ := strconv.Atoi(c.Query("page_number"))
	if currentPage <= 0 {
		currentPage = 1
	}
	itemsPerPage, _ := strconv.Atoi(c.Query("page_size"))
	if itemsPerPage <= 0 {
		itemsPerPage = 10
	}


    code, msg, data , totalItems := services.GetAllPlant(search, currentPage, itemsPerPage)
    // Atur headers pagination
    // c.Set("X-Total-Items", strconv.Itoa(totalItems))
    // c.Set("X-Items-Per-Page", strconv.Itoa(itemsPerPage))
    // c.Set("X-Current-Page", strconv.Itoa(currentPage))
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
			// Handle error jika terjadi kesalahan saat marshalling ke JSON
			panic(err)
		}

		// Atur header JSON ke dalam konteks Fiber
		c.Set("X-Pagination", string(jsonHeader))
    // Jika Anda memiliki lebih banyak informasi, Anda dapat menambahkannya di sini
    // Kembalikan respons JSON
    return c.Status(code).JSON(fiber.Map{
        "code":    code,
        "message": msg,
        "data":    data,
    })
}


func GetPlantById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	code, msg, data := services.GetPlantById(id)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

func StorePlant(c *fiber.Ctx) error {

	json := new(models.Plant)
	if err := c.BodyParser(json); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid JSON",
		})
	}
	code, msg := services.StorePlant(c.Locals("user_id").(int) ,json )
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func UpdatePlant(c *fiber.Ctx) error {
	json := new(models.Plant)
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
	code, msg := services.UpdatePlant(id, json)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func DeletePlant(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	code, msg := services.DeletePlant(id)
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}
