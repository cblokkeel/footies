package main

import (
	"net/http"
	"time"

	"github.com/cblokkeel/footies/api"
	"github.com/cblokkeel/footies/client"
)

func main() {
	baseClient := client.NewClient(&http.Client{Timeout: time.Second * 10})
	jsonPlaceHolderClient := client.NewJsonPlaceholderClient(baseClient)

	api := api.NewApi(jsonPlaceHolderClient)
	api.Start()
}
