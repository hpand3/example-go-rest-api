package handler

import (
	"gorestapi/src/contract"
	"gorestapi/src/domain"
	"gorestapi/src/mapper"
	"gorestapi/src/repository"
)

type CreateFeedItemCommand contract.ItemDto

type CreateFeedItemHandler struct {
	itemRepository repository.ItemRepository
}

func NewCreateFeedItemHandler(itemRepository repository.ItemRepository) *CreateFeedItemHandler {
	return &CreateFeedItemHandler{itemRepository}
}

func (h CreateFeedItemHandler) Handle(c CreateFeedItemCommand) {
	item := domain.NewItem(c.Title, c.Post)
	h.itemRepository.Add(*item)
}

type GetAllFeedsResponseDto struct {
	Items []contract.ItemDto `json:"items"`
}

type GetAllFeedsQueryHandler struct {
	itemRepository repository.ItemRepository
}

func NewGetAllFeedsQueryHandler(itemRepository repository.ItemRepository) *GetAllFeedsQueryHandler {
	return &GetAllFeedsQueryHandler{itemRepository}
}

func (h GetAllFeedsQueryHandler) Handle() GetAllFeedsResponseDto {
	items := h.itemRepository.GetAll()
	itemDtos := make([]contract.ItemDto, len(items))
	for index, item := range items {
		itemDtos[index] = mapper.MapItemToItemDto(item)
	}
	return GetAllFeedsResponseDto{
		Items: itemDtos,
	}
}
