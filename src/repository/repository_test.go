package repository

import (
	"gorestapi/src/domain"
	"testing"
)

func TestItemShouldExistWhenItemIsAddedToRepository(t *testing.T) {
	feedRepo := NewItemInMemoryRepository()

	feedRepo.Add(domain.Item{
		Title: "Hello world",
		Post:  "This is my first post",
	})

	if len(feedRepo.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestItemCanBeFetchedWhenItExistsInRepository(t *testing.T) {
	feedRepo := NewItemInMemoryRepository()

	feedRepo.Add(domain.Item{})

	if len(feedRepo.GetAll()) != 1 {
		t.Errorf("Not all items were fetched")
	}
}
