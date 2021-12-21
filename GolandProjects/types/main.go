package main

import "fmt"

// aggregate types(array, struct)
// struct is considered an aggregate type, because it can hold many pieces of information
// reference types(pointers, slices, maps, functions, channels)
// interface type

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XX2021",
		Year:          2021,
	}
	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
}
