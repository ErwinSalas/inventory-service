package services

import (
	"context"

	"github.com/ErwinSalas/inventory-service/models"
	inventorypb "github.com/ErwinSalas/inventory-service/proto"
	"github.com/jinzhu/gorm"
)

type InventoryService struct {
	inventorypb.UnimplementedInventoryServer
	datastore *gorm.DB
}

func NewInventoryService(datastore *gorm.DB) *InventoryService {
	return &InventoryService{
		datastore:                    datastore,
		UnimplementedInventoryServer: inventorypb.UnimplementedInventoryServer{},
	}
}

func (s *InventoryService) GetItem(ctx context.Context, in *inventorypb.ItemGetRequest) (*inventorypb.ItemGetResponse, error) {
	item := models.Item{}
	itemModel, err := item.FindItemByID(s.datastore, in.Id)

	if err != nil {
		return nil, err
	}
	return &inventorypb.ItemGetResponse{
		Item: &inventorypb.Item{
			Id:   itemModel.ID,
			Name: itemModel.Name,
		},
	}, nil
}

func (s *InventoryService) ListItems(ctx context.Context, in *inventorypb.ListItemsRequest) (*inventorypb.ListItemsResponse, error) {
	item := models.Item{}
	itemspbs := []*inventorypb.Item{}
	items, err := item.FindAllItems(s.datastore)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		itemspbs = append(itemspbs, &inventorypb.Item{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	return &inventorypb.ListItemsResponse{
		Items: itemspbs,
	}, nil
}
