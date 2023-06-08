package main

import (
	"github.com/dillonthompson/receipt-processor/api"
)

func main() {
	router := api.Router()
	router.Logger.Fatal(router.Start(":8080"))
}
