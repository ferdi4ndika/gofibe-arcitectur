package repositories

import (
	"fiber-ferdi/configs"
	"fiber-ferdi/models"
)

func GetAllModel(search string, page, pageSize int) ([]models.Model, int, error) {
    var Models []models.Model
    var totalItems int64

    db := configs.DB
    query := db.Model(&models.Model{})

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
    if err := query.Preload("User").Offset(offset).Limit(pageSize).Find(&Models).Error; err != nil {
        return nil, 0, err
    }

    return Models,int(totalItems), nil
}


func GetModelById(id int) (models.Model, error) {
	Model := models.Model{}
	db := configs.DB
	result := db.Model(&models.Model{}).Where("id = ?", id).Find(&Model)
	if result.Error != nil {
		return Model, result.Error
	}
	return Model, nil
}

func StoreModel(user_id int , json *models.Model) error {
	Model := models.Model{
		Information:  json.Information,
		Code:   json.Code,
		UserID:   user_id,	
	}
	db := configs.DB
	result := db.Create(&Model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateModel(id int, json *models.Model) error {
	Model := models.Model{
		Information:  json.Information,
		Code:   json.Code,
	}
	db := configs.DB
	result := db.Model(&models.Model{}).Where("id = ?", id).Updates(&Model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func DeleteModel(id int) error {
	Model := models.Model{}
	db := configs.DB
	result := db.Where("id = ?", id).Delete(&Model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
