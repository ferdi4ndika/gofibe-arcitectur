package repositories

import (
	"fiber-ferdi/configs"
	"fiber-ferdi/models"
)

func GetAllPart(search string, page, pageSize int) ([]models.Part, int, error) {
    var parts []models.Part
    var totalItems int64

    db := configs.DB
    query := db.Model(&models.Part{})

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
    if err := query.Preload("User").Offset(offset).Limit(pageSize).Find(&parts).Error; err != nil {
        return nil, 0, err
    }

    return parts,int(totalItems), nil
}

func GetAllPartWs() ([]models.Part, error) {
	part := []models.Part{}
	db := configs.DB
	result := db.Model(&models.Part{}).Preload("User").Find(&part)
	if result.Error != nil {
		return nil, result.Error
	}
	return part, nil
}
func GetExcle() ([]models.Part , error){
	part := []models.Part{}
	db := configs.DB
	result := db.Model(&models.Part{}).Find(&part)

	if result.Error !=nil{
		return nil , result.Error

	}
	return part ,nil
}

func GetPartById(id int) (models.Part, error) {
	part := models.Part{}
	db := configs.DB
	result := db.Model(&models.Part{}).Where("id = ?", id).Find(&part)
	if result.Error != nil {
		return part, result.Error
	}
	return part, nil
}

func StorePart(user_id int , json *models.Part) error {
	part := models.Part{
		Information:  json.Information,
		Code:   json.Code,
		UserID:   user_id,	
	}
	db := configs.DB
	result := db.Create(&part)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdatePart(id int, json *models.Part) error {
	part := models.Part{
		Information:  json.Information,
		Code:   json.Code,
	}
	db := configs.DB
	result := db.Model(&models.Part{}).Where("id = ?", id).Updates(&part)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func DeletePart(id int) error {
	part := models.Part{}
	db := configs.DB
	result := db.Where("id = ?", id).Delete(&part)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
