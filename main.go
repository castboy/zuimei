package main

import (
	"zuimei/core"
)

func main() {
	var person core.Animal = new(core.Person)
	person.Say()
	person.Eat("noodle")

	var bird core.Animal = new(core.Bird)
	bird.Say()
	bird.Eat("pet")
	core.Haha()

}
