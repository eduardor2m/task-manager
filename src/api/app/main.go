package main

import (
	"github.com/eduardor2m/task-manager/src/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(
		"./src/api/app/.env",
	)
	if err != nil {
		panic(err)
	}

	api := api.NewApi(&api.Options{})
	api.Serve()
}
