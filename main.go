package main

import (
	"fmt"

	dataLayer "wharleyinc.com/to-do/data"
	"wharleyinc.com/to-do/router"
)

func main() {

	dataLayer.InitDataWale()
	

	router.Router()

	fmt.Println("Server started on localhost:8080")
}
