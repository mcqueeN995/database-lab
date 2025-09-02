package main

import (
	"fmt"
	"mian/models"
	"mian/storage"
)

func main() {
	var namefile string
	fmt.Print("Enter your name file by open or create: \n")
	fmt.Scanln(&namefile)

	queue := storage.NewQueue(namefile)

	var command string
	for {
		fmt.Print("Enter your command: ")
		fmt.Scanln(&command)
		switch command {
		case "ADD":

			var name, surname, email, positon string
			var id int

			fmt.Print("Enter user id: ")
			fmt.Scanln(&id)
			fmt.Print("Enter user name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter user surname: ")
			fmt.Scanln(&surname)
			fmt.Print("Enter user email: ")
			fmt.Scanln(&email)
			fmt.Print("Enter positon: ")
			fmt.Scanln(&positon)

			user := models.User{
				ID:       id,
				Name:     name,
				Surname:  surname,
				Email:    email,
				Position: positon,
			}

			err := queue.Enqueue(user)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("User added\n")
			}
		case "DELETE":
			fmt.Print("Do you want to delete a user by %ID% or %Email%?\n")
			var answer string
			fmt.Scanln(&answer)
			if answer == "ID" || answer == "id" {
				var id int
				fmt.Print("Enter user id: ")
				fmt.Scanln(&id)
				err := queue.DeleteID(id)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("User deleted\n")
				}
			} else if answer == "Email" || answer == "email" {
				var email string
				fmt.Print("Enter email: ")
				fmt.Scanln(&email)
				err := queue.DeleteEmail(email)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("User deleted\n")
				}
			} else {
				fmt.Print("bro...I tell write ID or Email\n")
			}
		case "SEARCH":
			fmt.Print("Do you want to search a user by %ID% or %Email%?")
			var answer string
			fmt.Scanln(&answer)
			if answer == "ID" || answer == "id" {
				var id int
				fmt.Print("Enter user id: ")
				fmt.Scanln(&id)
				users, err := queue.SearchID(id)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("user ID: %d\n", users.ID)
					fmt.Printf("user Name: %s\n", users.Name)
					fmt.Printf("user Surname: %s\n", users.Surname)
					fmt.Printf("user Email: %s\n", users.Email)
					fmt.Printf("user Position: %s\n", users.Position)
				}
			} else if answer == "email" || answer == "Email" {
				var email string
				fmt.Print("Enter email: ")
				fmt.Scanln(&email)
				users, err := queue.SearchEmail(email)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("user ID: %d\n", users.ID)
					fmt.Printf("user Name: %s\n", users.Name)
					fmt.Printf("user Surname: %s\n", users.Surname)
					fmt.Printf("user Email: %s\n", users.Email)
					fmt.Printf("user Position: %s\n", users.Position)
				}
			} else {
				fmt.Print("bro...I tell write ID or Email\n")
			}
		case "UPDATE":
			fmt.Println("Enter user ID for update:")
			var id int
			fmt.Scanln(&id)
			var name, surname, email, positon string
			fmt.Print("Enter new user name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter new user surname: ")
			fmt.Scanln(&surname)
			fmt.Print("Enter new user email: ")
			fmt.Scanln(&email)
			fmt.Print("Enter new positon: ")
			fmt.Scanln(&positon)
			user := models.User{
				ID:       id,
				Name:     name,
				Surname:  surname,
				Email:    email,
				Position: positon,
			}
			err := queue.Update(id, user)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("User updated\n")
			}
		case "LIST":
			if queue.Len() == 0 {
				fmt.Println("No users found")
			} else {
				queue.PrintAll()
			}
		case "EXIT":
			return
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
