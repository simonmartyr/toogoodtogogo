# TooGoodToGoGo

TooGoodToGoGo is an unofficial API wrapper for the mobile application [too good to go](https://www.toogoodtogo.com)

## Getting Started 

### Install 

```shell
go get github.com/simonmartyr/toogoodtogogo
```

### Usage

```go
import "github.com/simonmartyr/toogoodtogogo"
```

creating a new client:

```go
client := toogoodtogo.New(
    &http.Client{},
    toogoodtogo.WithUsername("youremail@youremail.co.uk"),
)
```

It is also possible to create a client with an already know access tokens:

```go
client := toogoodtogo.New(
    &http.Client{},
    toogoodtogo.WithUsername("youremail@youremail.co.uk"),
    toogoodtogo.WithAuth("myId", "myAccessToken", "MyRefreshToken", "Mycookie"),
)
```

## Endpoints

This table highlights known endpoints and which are currently usable with the wrapper.

| Endpoint Name        | Description                                   | Supported |
|----------------------|-----------------------------------------------|-----------|
| Authenticate         | Login via magic/email link                    | &check;   |
| Refresh Auth         | Request new Access, Refresh tokens            | &check;   |
| Get Favorite Items   | Query for items marked as favorite            | &check;   |
| Search For Items     | Find items based on different search criteria | &check;   |
| Get Single Item      | Get details about one item by id              | &check;   |
| Set Item as Favorite | Mark an item as a favorite                    | &check;   |
| Create Order         | Begin process of buying item                  | &cross;   |
| Get Order Status     | Check current status of an order              | &cross;   |
| Cancel Order         | Cancel an order which isn't complete          | &cross;   |
| Signup               | Signup to too good to go                      | &cross;   |
| Get Active Orders    | Find currently active orders                  | &cross;   |
| Get Inactive Orders  | Find past orders                              | &cross;   |



