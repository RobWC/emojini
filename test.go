package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	fmt.Println("Hello Wolrd Emoji!")

	emoji.Println(":beer: Beer!!!")
	emoji.Println(":chicken: :poultry_leg: :chicken: :poultry_leg: :chicken: :nut_and_bolt: :poultry_leg: :chicken:")
	emoji.Println(":department_store:")
	emoji.Println(":angry:")
	emoji.Println(":rage:")
	emoji.Println(":u7a7a:")
	emoji.Println(":scream_cat:")
	emoji.Println(":radio_button:")

	pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
	fmt.Println(pizzaMessage)
}
