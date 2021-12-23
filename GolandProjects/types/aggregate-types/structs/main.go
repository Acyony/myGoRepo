package main

import "fmt"

type Car struct {
	NumOfTires int
	Luxury     bool
	BuckSeats  bool
	Make       string
	Model      string
	Year       int
}

func main() {
	// Aggregate Types - structs
	/*var myCar Car*/ // creating myCar var

	// populating myCar var
	/*myCar.NumOfTires = 4
	myCar.Luxury = false
	myCar.Make = "Volkswagen"*/

	// -------------=^_^=-----------
	//declaring and populating
	myCar := Car{
		NumOfTires: 4,
		Luxury:     true,
		BuckSeats:  true,
		Make:       "Volvo",
		Model:      "XC21",
		Year:       2019,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

}
