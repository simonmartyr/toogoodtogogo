package toogoodtogo

import (
	"net/http"
	"testing"
)

func TestClient_RefreshAuthentication(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/refresh_data.json")
	accessToken := c.accessToken
	refreshToken := c.refreshToken
	defer s.Close()
	err := c.RefreshAuthentication()
	if err != nil {
		t.Fatal(err)
	}
	if c.accessToken == accessToken {
		t.Error("Access token did not change")
	}
	if c.refreshToken == refreshToken {
		t.Error("refresh token did not change")
	}
}
