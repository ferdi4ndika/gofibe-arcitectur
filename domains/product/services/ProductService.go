package services

import (
	"fiber-ferdi/domains/product/repositories"
	"fiber-ferdi/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// func GetAllProduct() (int, string, []models.Product) {
// 	result, err := repositories.GetAllProduct()
// 	if err != nil {
// 		log.Fatal(err)
// 		return fiber.StatusBadRequest, err.Error(), nil
// 	}
// 	return fiber.StatusOK, "Get data product success", result
// }
func GetAllProduct() ([]map[string]interface{}, error) {
    products, err := repositories.GetAllProduct()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    var productsArray []map[string]interface{}

    for _, product := range products {
        total := product.Qty * product.Price
        productData := map[string]interface{}{
			//  'id'=>$this->id,
            // 'code' => $this->code,
            // 'information' => $this->name,
            // 'created_at' => $this->createdAt->format('Y-m-d H:i:s'),
            // 'updated_at' => $this->updatedAt->format('Y-m-d H:i:s'),
            "id":  product.Name,
            "code":  product.Name,
            "information":   product.Qty,
            "created_at": product.Price,
            "updated_at": total,
        }
        productsArray = append(productsArray, productData)
    }

    return productsArray, nil
}

func GetProductById(id int) (int, string, models.Product) {
	result, err := repositories.GetProductById(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error(), models.Product{}
	}
	return fiber.StatusOK, "Get data product success", result
}

func StoreProduct(json *models.Product) (int, string) {
	err := repositories.StoreProduct(json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Insert data product success"
}

func UpdateProduct(id int, json *models.Product) (int, string) {
	err := repositories.UpdateProduct(id, json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Update data product success"
}

func DeleteProduct(id int) (int, string) {
	err := repositories.DeleteProduct(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Delete data product success"
}
