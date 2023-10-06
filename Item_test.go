package toogoodtogo

import (
	"net/http"
	"testing"
)

func TestGetItems(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/get_items_data.json")
	defer s.Close()
	query := ItemsQueryFilter{SearchPhrase: "febo", Radius: 20, Latitude: 52.3538176190218, Longitude: 4.957415416727165}
	res, err := c.GetItems(&query)
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Error("failed to get items")
	}
}

func TestGetItem(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/get_item_data.json")
	defer s.Close()
	res, err := c.GetItem("1063134")
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Error("failed to get item")
	}
}
