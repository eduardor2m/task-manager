package main

import "github.com/eduardor2m/task-manager/src/api"

func main() {
	api := api.NewApi(&api.Options{})
	api.Serve()
}
