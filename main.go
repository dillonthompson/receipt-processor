package main

import (
	"net/http"

	"github.com/dillonthompson/receipt-processor/api"
)

func main() {
	api.RequestHandler()
	http.ListenAndServe(":8080", nil)
}
