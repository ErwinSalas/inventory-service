package seed

import (
	"log"

	"github.com/ErwinSalas/inventory-service/models"
	"github.com/jinzhu/gorm"
)

var items = []models.Item{
	models.Item{
		Name: "Macbook Pro",
	},
	models.Item{
		Name: "Ipad",
	},
	models.Item{
		Name: "Axe io",
	},
	models.Item{
		Name: "Boss Katana 50",
	},
	models.Item{
		Name: "Cargador de Mac",
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

	for i := range items {
		err = db.Debug().Model(&models.Item{}).Create(&items[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
