package mapper

import (
	"gorestapi/src/contract"
	"gorestapi/src/domain"
)

func MapItemToItemDto(item domain.Item) contract.ItemDto {
	return contract.ItemDto{
		Title: item.Title(),
		Post:  item.Post(),
	}
}
