package main

import (
	"github.com/eduardor2m/task-manager/src/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(
		"/Users/eduardomelo/Documents/projects/my-projects/task-manager/src/api/app/.env",
	)
	if err != nil {
		panic(err)
	}

	api := api.NewApi(&api.Options{})
	api.Serve()
}
