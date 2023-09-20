package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/muhammadsaefulr/inventori-barang/models"
)

func ConnectDatabase() (*gorm.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error loading env")
	}

	// DBusername := os.Getenv("DBusername")
	// DBpassword := os.Getenv("DBpassword")
	// DBurl := os.Getenv("DBurl")
	// DBport := os.Getenv("DBport")
	// DBname := os.Getenv("DBname")

	dsn := "root:@tcp(127.0.0.1:3306)/inventori?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.Barang{})

	return db, nil
}
