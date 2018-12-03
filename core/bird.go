package core

import "fmt"

type Bird struct {
}

func (b *Bird) Say() {
	fmt.Print("I am a bird")
}

func (b *Bird) Eat(food string) {
	fmt.Println("\n bird eat %s\n", food)
}
