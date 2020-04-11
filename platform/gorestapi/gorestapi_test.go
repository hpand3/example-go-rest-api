package gorestapi

import "testing"

func TestAdd(t *testing.T) {
	api := New()

	api.Add(Item{
		Title: "Hello world",
		Post: "This is my first post",
	})

	if len(api.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	api := New()

	api.Add(Item{})

	if len(api.GetAll()) != 1 {
		t.Errorf("Not all items were fetched")
	}
}