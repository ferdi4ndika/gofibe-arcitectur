package repositories

import (
	"fiber-ferdi/configs"
	"fiber-ferdi/models"
	"fmt"
)

// func GetAllPlant(search string) ([]models.Plant, error) {
//     var plants []models.Plant
//     db := configs.DB
//     query := db.Model(&models.Plant{})

//     // Jika ada kata kunci pencarian, tambahkan kondisi WHERE ke query
//     if search != "" {
//         query = query.Where("information LIKE ?", "%"+search+"%")
//     }

//     // Jalankan query dan ambil hasilnya
//     if err := query.Preload("User").Find(&plants).Error; err != nil {
//         return nil, err
//     }

//     return plants, nil
// }
func datata() {
    // Panggil fungsi GetAllPlant untuk dijalankan
   fmt.Println("GetAllPlant function implementation")
}
func GetAllPlant(search string, page, pageSize int) ([]models.Plant, int, error) {
    var plants []models.Plant
    var totalItems int64

    db := configs.DB
    query := db.Model(&models.Plant{})

    // Jika ada kata kunci pencarian, tambahkan kondisi WHERE ke query
    if search != "" {
        query = query.Where("information LIKE ?", "%"+search+"%")
    }

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

    // Menghitung offset berdasarkan halaman dan jumlah item per halaman
    offset := (page - 1) * pageSize

    // Jalankan query dengan limit dan offset
    if err := query.Preload("User").Offset(offset).Limit(pageSize).Find(&plants).Error; err != nil {
        return nil, 0, err
    }

    return plants,int(totalItems), nil
}

func GetAllPlantWs() ([]models.Plant, error) {
	plant := []models.Plant{}
	db := configs.DB
	result := db.Raw("SELECT code FROM plants WHERE id = ?", 1).Scan(&plant)//Query manual
	// result := db.Model(&models.Plant{}).Where("id = ?", 1).Find(&plant)
	if result.Error != nil {
		return nil, result.Error
	}
	return plant, nil
}

func GetPlantById(id int) (models.Plant, error) {
	plant := models.Plant{}
	db := configs.DB
	result := db.Model(&models.Plant{}).Where("id = ?", id).Find(&plant)
	if result.Error != nil {
		return plant, result.Error
	}
	return plant, nil
}

func StorePlant(user_id int , json *models.Plant) error {
	plant := models.Plant{
		Information:  json.Information,
		Code:   json.Code,
		UserID:   user_id,	
	}
	db := configs.DB
	result := db.Create(&plant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdatePlant(id int, json *models.Plant) error {
	plant := models.Plant{
		Information:  json.Information,
		Code:   json.Code,
	}
	db := configs.DB
	result := db.Model(&models.Plant{}).Where("id = ?", id).Updates(&plant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func DeletePlant(id int) error {
	plant := models.Plant{}
	db := configs.DB
	result := db.Where("id = ?", id).Delete(&plant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
