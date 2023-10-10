package main

import (
	"fmt"
	"github.com/simonmartyr/toogoodtogogo"
	"net/http"
)

func main() {
	client := toogoodtogo.New(
		&http.Client{},
		//change this
		toogoodtogo.WithUsername("youremail@youremail.co.uk"),
	)
	err := client.Authenticate()
	if err != nil {
		fmt.Println("Failed to auth")
		return
	}
	credentials := client.GetCredentials()
	fmt.Println("Successful auth")
	fmt.Println(fmt.Sprintf("userId: %s", credentials.UserId))
	fmt.Println(fmt.Sprintf("Access token: %s", credentials.AccessToken))
	fmt.Println(fmt.Sprintf("Refresh token: %s", credentials.RefreshToken))
	fmt.Println(fmt.Sprintf("cookie: %s", credentials.Cookie))
}
