package repositories

import (
	"fiber-ferdi/configs"
	"fiber-ferdi/models"
)

func GetAllOp(search string, page, pageSize int) ([]models.Op, int, error) {
    var Ops []models.Op
    var totalItems int64

    db := configs.DB
    query := db.Model(&models.Op{})

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
    if err := query.Preload("User").Offset(offset).Limit(pageSize).Find(&Ops).Error; err != nil {
        return nil, 0, err
    }

    return Ops,int(totalItems), nil
}


func GetOpById(id int) (models.Op, error) {
	Op := models.Op{}
	db := configs.DB
	result := db.Model(&models.Op{}).Where("id = ?", id).Find(&Op)
	if result.Error != nil {
		return Op, result.Error
	}
	return Op, nil
}

func StoreOp(user_id int , json *models.Op) error {
	Op := models.Op{
		Information:  json.Information,
		Code:   json.Code,
		UserID:   user_id,	
	}
	db := configs.DB
	result := db.Create(&Op)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateOp(id int, json *models.Op) error {
	Op := models.Op{
		Information:  json.Information,
		Code:   json.Code,
	}
	db := configs.DB
	result := db.Model(&models.Op{}).Where("id = ?", id).Updates(&Op)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func DeleteOp(id int) error {
	Op := models.Op{}
	db := configs.DB
	result := db.Where("id = ?", id).Delete(&Op)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
