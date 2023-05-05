package main

import (
	"oauth-poc/api"
	"oauth-poc/infra"
)

func main() {
	api := api.NewAPI(infra.NewGinRouter())
	api.Start(":8081")
}
