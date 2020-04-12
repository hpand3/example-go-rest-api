package handler

import (
	"gorestapi/src/domain"
	"gorestapi/src/repository"
)

type CreateFeedItemCommand struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type CreateFeedItemHandler struct {
	itemRepository repository.ItemRepository
}

func NewCreateFeedItemHandler(itemRepository repository.ItemRepository) *CreateFeedItemHandler {
	return &CreateFeedItemHandler{itemRepository}
}

func (h CreateFeedItemHandler) Handle(c CreateFeedItemCommand) {
	item := domain.Item{
		Title: c.Title,
		Post:  c.Post,
	}
	h.itemRepository.Add(item)
}

type GetAllFeedsQuery struct{}
type GetAllFeedsResponse struct {
	Items []domain.Item `json:"items"`
}

type GetAllFeedsQueryHandler struct {
	itemRepository repository.ItemRepository
}

func NewGetAllFeedsQueryHandler(itemRepository repository.ItemRepository) *GetAllFeedsQueryHandler {
	return &GetAllFeedsQueryHandler{itemRepository}
}

func (h GetAllFeedsQueryHandler) Handle(q GetAllFeedsQuery) GetAllFeedsResponse {
	return GetAllFeedsResponse{
		Items: h.itemRepository.GetAll(),
	}
}
