package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"nickname"`
}

func (u *Item) SaveItem(db *gorm.DB) (*Item, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &Item{}, err
	}
	return u, nil
}

func (u *Item) FindAllItems(db *gorm.DB) ([]Item, error) {
	var err error
	Items := []Item{}
	err = db.Debug().Model(&Item{}).Limit(100).Find(&Items).Error
	if err != nil {
		return []Item{}, err
	}
	return Items, err
}

func (u *Item) FindItemByID(db *gorm.DB, uid uint64) (*Item, error) {
	err := db.Debug().Model(Item{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Item{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Item{}, errors.New("Item Not Found")
	}
	return u, err
}

func (u *Item) UpdateAItem(db *gorm.DB, uid uint64) (*Item, error) {

	db = db.Debug().Model(&Item{}).Where("id = ?", uid).Take(&Item{}).UpdateColumns(
		map[string]interface{}{
			"Name": u.Name,
		},
	)
	if db.Error != nil {
		return &Item{}, db.Error
	}
	// This is the display the updated Item
	err := db.Debug().Model(&Item{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Item{}, err
	}
	return u, nil
}

func (u *Item) DeleteAItem(db *gorm.DB, uid uint64) (int64, error) {

	db = db.Debug().Model(&Item{}).Where("id = ?", uid).Take(&Item{}).Delete(&Item{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
