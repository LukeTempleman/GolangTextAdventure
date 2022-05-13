package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

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
}

func randFirearm() {
	c := colly.NewCollector()

	c.OnHTML("div", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://perchance.org/pi96s9irte")

}

func randMelee() {

}

func randSecret() {

}

func randBodyPart() {
	link := "https://perchance.org/body-part"
}

func randLocation() {

}

func randEnemy() {
	link := "https://perchance.org/monster-type"
}
