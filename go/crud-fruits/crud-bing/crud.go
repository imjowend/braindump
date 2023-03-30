package main

import (
	"fmt"
)

type Fruit struct {
	Name string
}

var fruits []Fruit

func createFruit(name string) {
	fruits = append(fruits, Fruit{Name: name})
}

func readFruits() {
	fmt.Println("Fruits:")
	for _, fruit := range fruits {
		fmt.Println(fruit.Name)
	}
}

func updateFruit(oldName string, newName string) {
	for i := 0; i < len(fruits); i++ {
		if fruits[i].Name == oldName {
			fruits[i].Name = newName
			break
		}
	}
}

func deleteFruit(name string) {
	for i := 0; i < len(fruits); i++ {
		if fruits[i].Name == name {
			fruits = append(fruits[:i], fruits[i+1:]...)
			break
		}
	}
}

func main() {
	createFruit("Apple")
	createFruit("Banana")
	readFruits()
	updateFruit("Banana", "Orange")
	readFruits()
	deleteFruit("Apple")
	readFruits()
}
