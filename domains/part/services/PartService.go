package services

import (
	"fiber-ferdi/domains/part/repositories"
	"fiber-ferdi/models"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tealeg/xlsx"
)

func GetAllPart(search string, currentPage, itemsPerPage int) (int, string, []models.Part, int) {
    // Mendapatkan data part dari repository dengan menggunakan search, currentPage, dan itemsPerPage
    result, totalCount, err := repositories.GetAllPart(search, currentPage, itemsPerPage)
    if err != nil {
        log.Fatal(err)
        return fiber.StatusBadRequest, err.Error(), nil, 0
    }
    return fiber.StatusOK, "Get data part success", result, totalCount
}



func GetAllPartWs() ([]map[string]interface{}, error) {
    parts, err := repositories.GetAllPartWs()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    var partsArray []map[string]interface{}

    for _, part := range parts {
        // total := part.Qty * part.Price
        partData := map[string]interface{}{
            "id":  part.ID,
            "code":  part.Code,
            "information":   part.Information,
            "created_at": part.CreatedAt,
            "updated_at":part.CreatedAt,
        }
        partsArray = append(partsArray, partData)
    }

    return partsArray, nil
}

func GetExcle(c *fiber.Ctx) error {
    // Ambil data dari repository
    parts, err := repositories.GetExcle()
    if err != nil {
        // Mengembalikan kesalahan tanpa menggunakan log.Fatal
        return err
    }

    // Buat file Excel baru dengan nama yang unik
    filename := fmt.Sprintf("excel_%d.xlsx", time.Now().UnixNano())
    file := xlsx.NewFile()
    sheet, err := file.AddSheet("Data")
    if err != nil {
        // Mengembalikan kesalahan tanpa menggunakan log.Fatal
        return err
    }

    // Tambahkan header ke file Excel
    headers := []string{"ID", "Code", "Information", "Created At", "Updated At"}
    headerRow := sheet.AddRow()
    for _, header := range headers {
        cell := headerRow.AddCell()
        cell.SetString(header)
    }

    // Tambahkan data ke file Excel
    for _, part := range parts {
        row := sheet.AddRow()
        row.AddCell().SetInt(part.ID)
        row.AddCell().SetString(part.Code)
        row.AddCell().SetString(part.Information)
        row.AddCell().SetString(part.CreatedAt.String()) // Sesuaikan dengan format yang Anda inginkan
        row.AddCell().SetString(part.UpdatedAt.String()) // Sesuaikan dengan format yang Anda inginkan
    }

    // // Simpan file Excel dengan nama yang unik
    // err = file.Save(filename)
    // if err != nil {
    //     // Mengembalikan kesalahan tanpa menggunakan log.Fatal
    //     return err
    // }

    // defer func() {
    //     // Hapus file Excel sementara setelah selesai
    //     err := os.Remove(filename)
    //     if err != nil {
    //         log.Println("Failed to delete temporary file:", err)
    //     }
    // }()

    // // Kirim file Excel sebagai respons
    // return c.SendFile(filename)

	/// TAMAPA MENYIPAN FILE NYA KE DIRETORY

	//   c.Set("Content-Disposition", "attachment; filename=yourfilename.xlsx")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
    c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

    // Tulis file Excel langsung ke http.ResponseWriter
    err = file.Write(c.Response().BodyWriter())
    if err != nil {
        return err
    }

    // Selesai menulis, kembalikan status sukses
    return nil
}



func GetPartById(id int) (int, string, models.Part) {
	result, err := repositories.GetPartById(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error(), models.Part{}
	}
	return fiber.StatusOK, "Get data part success", result
}

func StorePart(user_id int ,json *models.Part) (int, string) {
	err := repositories.StorePart(user_id ,json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Insert data part success"
}

func UpdatePart(id int, json *models.Part) (int, string) {
	err := repositories.UpdatePart(id, json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Update data part success"
}

func DeletePart(id int) (int, string) {
	err := repositories.DeletePart(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Delete data part success"
}
