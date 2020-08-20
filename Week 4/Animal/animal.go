package main

import (
"errors"
"fmt"
	"strings"
)

type Named interface {
	getName() string
}

type Animal interface {
	Named
	Eat()
	Move()
	Speak()
}


type Cow struct {name string}
func (c Cow) getName() string {return c.name}
func (c Cow) Eat() {fmt.Println("The Cow", c.name, "eats grass")}
func (c Cow) Move() {fmt.Println("The Cow", c.name, "walk")}
func (c Cow) Speak() {fmt.Println("The Cow", c.name, "sounds like moo")}

type Bird struct {name string}
func (b Bird) getName() string {return b.name}
func (b Bird) Eat() {fmt.Println("The Bird", b.name, "eats worms")}
func (b Bird) Move() {fmt.Println("The Bird", b.name, "fly")}
func (b Bird) Speak() {fmt.Println("The Bird", b.name, "sounds like peep")}

type Snake struct {name string}
func (s Snake) getName() string {return s.name}
func (s Snake) Eat() {fmt.Println("The Snake", s.name, "eats mice")}
func (s Snake) Move() {fmt.Println("The Snake", s.name, "slither")}
func (s Snake) Speak() {fmt.Println("The Snake", s.name, "sounds like hsss")}



func isValidCommand(command string) bool {return command == "newanimal" || command == "query"}

func isValidAnimal(animal string) bool {return animal == "cow" || animal == "bird" || animal == "snake"}

func isValidAnimalProperty(property string) bool {return property == "eat" || property == "move" || property == "speak"}

func getRequest() (string, string, string) {
	var command, name, option string
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&command, &name, &option); err != nil {
			fmt.Println("Error:", err)
		} else {
			command = strings.ToLower(command)
			if isValidCommand(command) {
				if command == "newanimal" {
					option = strings.ToLower(option)
					if isValidAnimal(option) {
						break
					} else {
						fmt.Println("Invalid animal type. Please use 'cow', 'bird' or 'snake'.")
					}
				} else if command == "query" {
					if isValidAnimalProperty(option) {
						break
					} else {
						fmt.Println("Invalid animal property. Please use 'eat', 'move' or 'speak'.")
					}
				}
			} else {
				fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
			}
		}
	}
	return command, name, option
}

func appendAnimal(animals *[]Animal, animal Animal) {
	*animals = append(*animals, animal)
	fmt.Println("Created it!")
}

func findAnimal(animals []Animal, name string) (Animal, error) {
	var foundAnimal Animal
	var err = errors.New("Animal with that name not found")
	for _, animal := range animals {
		if Named(animal).getName() == name {
			foundAnimal = animal
			err = nil
			break
		}
	}
	return foundAnimal, err
}



func main() {
	var animals []Animal

	for {
		command, name, option := getRequest()
		switch command {
		case "newanimal":
			switch option {
			case "cow":
				appendAnimal(&animals, Cow{name})
			case "bird":
				appendAnimal(&animals, Bird{name})
			case "snake":
				appendAnimal(&animals, Snake{name})
			}
		case "query":
			var animal Animal
			var err error
			if animal, err = findAnimal(animals, name); err != nil {
				fmt.Println(err)
				break
			}
			switch option {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		}
	}
}