package main

import "fmt"

func main() {
	homeScreen()
}

func homeScreen() {
	fmt.Println("\n\n\n\n\n\n\n28East Zombie Survival!")
	fmt.Println("-----------------------------")
	fmt.Println("Please Select an option")
	fmt.Println("1) Start Game")
	fmt.Println("2) How To Play")
	fmt.Println("3) Exit")

	var answer string

	for answer != "1" && answer != "2" && answer != "3" {
		fmt.Print(":")
		fmt.Scan(&answer)

		switch answer {
		case "1":
			fmt.Print("yeet")
		case "2":
			tutorial()
		case "3":
			fmt.Print("Goodbye")
		default:
			fmt.Println("Not an answer... Try again")
		}

	}

}


