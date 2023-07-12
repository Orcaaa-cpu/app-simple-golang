package main

import (
	"log"

	"product/routes"
)

func main() {
	e := routes.Init()

	err := e.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
