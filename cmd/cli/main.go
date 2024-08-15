package main

import (
	"fmt"
	"os"

	"Github.com/julianob-dev/ejercicio-2/internal/application"
	"Github.com/julianob-dev/ejercicio-2/internal/domain/interfaces"
	"Github.com/julianob-dev/ejercicio-2/internal/repository/memory"
)


func main() {
	
	args := os.Args
	key := args[1]
	var repo interfaces.Repository
	repo, err := memory.New()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := application.New(repo)
	sector, err := app.GetById(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("se encontro el sector ",sector)
}