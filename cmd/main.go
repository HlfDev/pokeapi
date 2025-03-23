package main

import "github.com/HlfDev/pokeapi-wrapper/internal/bootstrap"

// @title           Pokemon API
// @version         1.0
// @description     A simple REST API for Pokemon data
// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	bootstrap.Initialize()
}
