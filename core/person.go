package core

import "fmt"

type Person struct{}

func (p *Person) Say() {
	fmt.Print("I am one person")
}

func (p *Person) Eat(food string) {
	fmt.Printf("\n I eat %s\n", food)
}
