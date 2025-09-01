package main

import (
	"fmt"
	"mian/models"
	"mian/storage"
)

func main() {
	var namefile string
	fmt.Print("Enter your name file: \n")
	fmt.Scanln(&namefile)

	queue := storage.NewQueue(namefile)

	var command string
	for {
		switch command {
		case "":
		}
	}

}

func printUser(user models.User) {
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Surname: %s\n", user.Surname)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Position: %s\n", user.Position)
}
