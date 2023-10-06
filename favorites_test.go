package toogoodtogo

import (
	"net/http"
	"testing"
)

func TestClient_GetFavorites(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/favorite_data.json")
	defer s.Close()
	res, err := c.GetFavorites(52.3538176190218, 4.957415416727165, 500, 0, 50)
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Error("failed to get favorites")
	}
}
