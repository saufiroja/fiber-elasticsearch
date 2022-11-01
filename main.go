package main

import "elasticsearch/fiber-elasticsearch/infrastructure/server"

func main() {
	app := server.Server()

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
