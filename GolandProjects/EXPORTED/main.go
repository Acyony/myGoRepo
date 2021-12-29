package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "Alcione", LastName: "Ribeiro", Salary: 60000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 40000, FullTime: false},
	{FirstName: "John", LastName: "Travolta", Salary: 70000, FullTime: false},
	{FirstName: "Mickey", LastName: "Mouse", Salary: 40000, FullTime: true},
	{FirstName: "Mark", LastName: "Smith", Salary: 85000, FullTime: false},
}

func main() {
	// Initializing the package staff
	myStaff := staff.Office{
		AllStaff: employees,
	}

	//log.Println(myStaff.All())
	staff.OverPaidLimit = 60000
	log.Println("Overpaid staff", myStaff.Overpaid())
	log.Println("UnderPaid staff", myStaff.UnderPaid())

	//myStaff.notVisible() => not an exported function. Not possible to access here.
}
