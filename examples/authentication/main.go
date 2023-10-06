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
	access, refresh, cookie := client.GetCredentials()
	fmt.Println("Successful auth")
	fmt.Println(fmt.Sprintf("Access token: %s", access))
	fmt.Println(fmt.Sprintf("Refresh token: %s", refresh))
	fmt.Println(fmt.Sprintf("cookie: %s", cookie))
}
