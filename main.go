package main

import (
	"fmt"

	dataLayer "wharleyinc.com/to-do/data"
	"wharleyinc.com/to-do/router"
)

func main() {

	dataLayer.UseMongoDbAtlas()

	router.Router()

	fmt.Println("Server started on localhost:8080")
}
