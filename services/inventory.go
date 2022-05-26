package services

import (
	"context"
	"math/rand"

	"github.com/ErwinSalas/inventory-service/models"
	inventorypb "github.com/ErwinSalas/inventory-service/proto/inventory"
	"github.com/jinzhu/gorm"
)

type InventoryService struct {
	inventorypb.UnimplementedInventoryServer
	datastore *gorm.DB
}

func NewInventoryService(datastore *gorm.DB) *InventoryService {
	return &InventoryService{
		datastore:                    datastore,
		UnimplementedInventoryServer: &UnimplementedInventoryServer{},
	}
}

func (s *InventoryService) GetItem(ctx context.Context, in *inventorypb.GetItemRequest) (*inventorypb.GetItemResponse, error) {
	id := make([]byte, 4)
	rand.Read(id)
	return &inventorypb.GetItemResponse{
		item: &inventorypb.Item{
			id:   id,
			name: "Test item",
		},
	}, nil
}

func (s *InventoryService) List(ctx context.Context, in *inventorypb.GetItemRequest) (*inventorypb.GetItemResponse, error) {
	item := models.Item{}
	itemspbs := []*inventorypb.Item{}
	items, err := item.FindAllItems(InventoryService.datastore)
	if err != nil {
		return nil, err
	}

	for items := range items {
		append(itemspbs, &inventorypb.Item{
			id:   item.ID,
			name: item.Name,
		})
	}
}
