package service

import (
	"github.com/Termpao/repository"
)

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) CreateItem(data ItemRes) error {
	item := repository.ItemOrder{
		ItemName:        data.ItemName,
		ItemDescription: data.ItemDescription,
		ItemPrice:       data.ItemPrice,
		ProductId:       data.ItemProdctId,
	}

	if err := s.repo.Create(item); err != nil {
		return err
	}

	return nil
}

func (s *itemService) DeleteItem(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *itemService) EditItem(id int, data ItemRes) error {

	item := repository.ItemOrder{
		ItemName:        data.ItemName,
		ItemDescription: data.ItemDescription,
		ItemPrice:       data.ItemPrice,
	}
	if err := s.repo.Edit(id, item); err != nil {
		return err
	}

	return nil
}

func (s *itemService) GetItems() ([]ItemRes, error) {
	items, err := s.repo.Gets()
	if err != nil {
		return nil, err
	}

	var newitems []ItemRes
	for _, item := range items {
		newitems = append(newitems, ItemRes{
			ItemID:          item.ItemID,
			ItemName:        item.ItemName,
			ItemPrice:       item.ItemPrice,
			ItemDescription: item.ItemDescription,
		})

	}

	return newitems, nil
}

func (s *itemService) GetItem(id int) (*ItemRes, error) {
	item, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return &ItemRes{

		ItemID:          item.ItemID,
		ItemName:        item.ItemName,
		ItemPrice:       item.ItemPrice,
		ItemDescription: item.ItemDescription,
	}, nil

}

func (s *itemService) GetAllItemFromProduct(data ItemRes) ([]ItemRes, error) {

	items, err := s.repo.GetsItemFromProduct(repository.ItemOrder{ProductId: data.ItemProdctId})

	if err != nil {
		return nil, nil
	}

	var itemRes []ItemRes

	for _, item := range items {
		itemRes = append(itemRes, ItemRes{
			ItemID:          item.ItemID,
			ItemName:        item.ItemName,
			ItemPrice:       item.ItemPrice,
			ItemDescription: item.ItemDescription,
		})
	}
	return itemRes, nil

}
