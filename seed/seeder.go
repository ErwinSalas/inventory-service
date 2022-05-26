package seed

import (
	"log"

	"github.com/ErwinSalas/inventory-service/models"
	"github.com/jinzhu/gorm"
)

var items = []models.Item{
	models.Item{
		Name: "Steven victor",
	},
	models.Item{
		Name: "Martin Luther",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Item{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Item{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range items {
		err = db.Debug().Model(&models.Items{}).Create(&items[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
