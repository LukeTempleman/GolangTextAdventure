package main

import "fmt"

func tutorialStart() {
	fmt.Println("Welcome to the tutorial")
	fmt.Println("Here you will learn the basics of fighting the undead")
	fmt.Println("To get started , a random weapon will spawn")
	fmt.Println("You may choose the catagory of weapon, but not the actual weapon itself...")
	fmt.Println("1) Melee Weapons")
	fmt.Println("2) Firearms ")
	fmt.Println("3) Secret Weapons....")

	var answer string
	var memeChoice string

	for answer != "1" && answer != "2" && answer != "3" {
		fmt.Print(":")
		fmt.Scan(&answer)

		switch answer {
		case "1":
			randFirearm()
		case "2":
			randFirearm()
		case "3":
			fmt.Print("Are you sure?? (y/n)")
			fmt.Scan(&memeChoice)
			if memeChoice == "y" {
				fmt.Println("Mwahahaha wise choice")
			} else {
				fmt.Println("too bad.. lets begin")
			}

		default:
			fmt.Println("Not an answer... Try again")
		}

	}

}

func levelStart() {
	fmt.Println("too bad.. lets begin")
	fmt.Println("too bad.. lets begin")

}
