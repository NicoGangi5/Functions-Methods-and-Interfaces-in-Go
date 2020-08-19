package main

import (
	bu "bufio"
	"fmt"
	"os"
	stg "strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (an Animal) Eat() {
	fmt.Println("Your animal eats", an.food)
}
func (an Animal) Move()  {
	fmt.Println("Your animal", an.locomotion)
}
func (an Animal) Speak()  {
	fmt.Println("Your animal sounds like", an.noise)
}

func main()  {
	var cow, bird, snake Animal

	cow.food =  "grass"
	cow.locomotion = "walk"
	cow.noise = "moo"

	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"

	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hsss"

	for {
		reader := bu.NewReader(os.Stdin)

		fmt.Print(">")

		input, _ := reader.ReadString('\n')
		word := stg.TrimSuffix(input, "\n")
		word = stg.ToLower(word)

		result :=  stg.Split(word, " ")
		word1 := result [0]
		word2 := result [1]

		switch word1 {
		case "cow":
			switch word2 {
			case "eat": cow.Eat()
			case "move": cow.Move()
			case "speak": cow.Speak()
			default: fmt.Println("The action doesn't exist")
			}
		case "bird":
			switch word2 {
			case "eat": bird.Eat()
			case "move": bird.Move()
			case "speak": bird.Speak()
			default: fmt.Println("The action doesn't exist")
			}
		case "snake":
			switch word2 {
			case "eat": snake.Eat()
			case "move": snake.Move()
			case "speak": snake.Speak()
			default: fmt.Println("The action doesn't exist")
			}
		default:
			fmt.Println("The animal doesn't exist")
		}
	}
}
