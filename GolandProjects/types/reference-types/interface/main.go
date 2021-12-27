package main

import "fmt"

// interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

// Dog is the type for dogs
type Dog struct {
	Name      string
	Sound     string
	NumOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumOfLegs
}

// Cat is the type of cats
type Cat struct {
	Name      string
	Sound     string
	NumOfLegs int
	HasTail   bool
}

func (d *Cat) Says() string {
	return d.Sound
}

func (d *Cat) HowManyLegs() int {
	return d.NumOfLegs
}

func main() {
	// ask a riddle
	dog := Dog{
		Name:      "dog",
		Sound:     "woof",
		NumOfLegs: 4,
	}
	Riddle(&dog)
	/*-----------------=^.^=-----------------------*/
	var cat Cat
	cat.Name = "cat"
	cat.NumOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true
	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
