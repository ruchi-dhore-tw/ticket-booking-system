package app

import "fmt"

func StartServer() {
	fmt.Println("Starting server")
	router := addRoutes()
	router.Run(":8080")
}
