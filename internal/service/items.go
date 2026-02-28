package service

import (
	"context"
	"errors"
)

type ItemDb interface {
	Create(ctx context.Context, data CreateParams) (int, error)
	Update(ctx context.Context, id int, params UpdateParams) error
	Get(ctx context.Context, id int) (Item, error)
	Delete(ctx context.Context, id int) (int, error)
}

type ItemService struct {
	repo ItemDb
}

func NewItemService(db ItemDb) *ItemService {
	return &ItemService{db}
}

func (s *ItemService) CreateItem(ctx context.Context, params CreateParams) (int, error) {
	var id int
	var err error

	if id, err = s.repo.Create(ctx, params); err != nil {
		return id, errors.New("Create item error")
	}

	return id, err
}

func (s *ItemService) UpdateItem(ctx context.Context, id int, params UpdateParams) error {
	var err error

	if err = s.repo.Update(ctx, id, params); err != nil {
		return errors.New("Update item error")
	}

	return err
}

func (s *ItemService) GetItem(ctx context.Context, id int) (Item, error) {
	var item Item
	var err error

	if item, err = s.repo.Get(ctx, id); err != nil {
		return item, errors.New("Get item error")
	}

	return item, err
}

func (s *ItemService) DeleteItem(ctx context.Context, id int) (int, error) {
	var dId int
	var err error

	if dId, err = s.repo.Delete(ctx, id); err != nil {
		return dId, errors.New("Del item error")
	}

	return dId, err
}
