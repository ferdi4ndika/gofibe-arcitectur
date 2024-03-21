package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDB() {
	godotenv.Load()
	var err error
	dsn := os.Getenv("DATA_BASES")
	// dsn := "root:@tcp(127.0.0.1:3306)/fiber_ferdi?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

// package configs

// import (
// 	"fiber-ferdi/models"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var (
//     db *gorm.DB
// )

// func ConnDB() {
//     // Koneksikan ke database
// godotenv.Load()
//     var err error
//     dsn := os.Getenv("DATA_BASES")
//     db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//     if err != nil {
//         panic("Failed to connect to database!")
//     }

//     // Migrasi otomatis
//     err = db.AutoMigrate( &models.Op{},&models.Model{},&models.Part{},&models.Role{}, &models.User{}, &models.Product{}, &models.Plant{})
//     if err != nil {
//         panic("Failed to perform auto migration!")
//     }
// }

// func GetDB() *gorm.DB {
//     return db
// }







