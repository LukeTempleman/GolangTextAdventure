package main

import "fmt"

func takeDamage(currenthealth int, damageTaken int) int {
	healthFinal := currenthealth - damageTaken
	if healthFinal >= 0 {
		gameOver()
	}
	return healthFinal
}

func gameOver() {
	fmt.Println("Welcome to the tutorial")
	fmt.Println("Welcome to the tutorial")
	fmt.Println("Welcome to the tutorial")
}
