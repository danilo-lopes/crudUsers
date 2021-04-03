package main

import (
	"fmt"
	"os"

	"github.com/danilo-lopes/crudUsers/services"
)

// const (
// 	userFile string = "/etc/passwd"
// )

func main() {

	for {
		fmt.Println("Software Version: 1.0")

		fmt.Println("\n 1 - Add Users \n 2 - Lock/Unlock Users \n 3 - Update an User \n 4 - Exit")

		var command int
		fmt.Scan(&command)

		switch command {
		case 1:
			// Add users
			fmt.Println("Adding users..")
			services.RegisterUsers()

		case 2:
			// Lock/Unlock users
			fmt.Println("do something else..")

		case 3:
			// Update user
			fmt.Println("Locking user..")

		case 4:
			fmt.Println("Closing..")
			os.Exit(0)

		case 5:
			fmt.Println("Option Not Found")
			os.Exit(-1)
		}
	}
}
