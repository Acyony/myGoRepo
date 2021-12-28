package main

import "fmt"

type Vehicle struct {
	NumOfWHeels     int
	NumOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	// the Car now has access to Vehicle
	Vehicle Vehicle
}

func (vehicle Vehicle) ShowDetails() {
	fmt.Println("Number of passengers:", vehicle.NumOfPassengers)
	fmt.Println("Number of wheels:", vehicle.NumOfWHeels)
}

func (car Car) ShowMyCar() {
	fmt.Println("Make:", car.Make)
	fmt.Println("Model:", car.Model)
	fmt.Println("Year:", car.Year)
	fmt.Println("Is Electric:", car.IsElectric)
	fmt.Println("Is Hybrid:", car.IsHybrid)

	car.Vehicle.ShowDetails()
}

func main() {
	suv := Vehicle{
		NumOfWHeels:     4,
		NumOfPassengers: 6,
	}

	volvoX2021 := Car{
		Make:       "Volvo",
		Model:      "XD2022",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle:    suv,
	}

	volvoX2021.ShowMyCar()
	fmt.Println()

	teslaXX := Car{
		Make:       "Tesla",
		Model:      "Model XX",
		Year:       2023,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}
	teslaXX.Vehicle.NumOfPassengers = 7
	teslaXX.ShowMyCar()
}
