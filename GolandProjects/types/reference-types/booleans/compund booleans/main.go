package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	Jack := Employee{
		Name:     "Jack Smith",
		Age:      30,
		Salary:   60000,
		FullTime: false,
	}

	Jones := Employee{
		Name:     "Jill Jones",
		Age:      35,
		Salary:   65000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, Jack)
	employees = append(employees, Jones)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 or older")
		} else {
			fmt.Println(x.Name, "is under 30")

		}
		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 5000 and  is over than 30")
		} else {
			fmt.Println("Either", x.Name, "makes less than 5000 or is under than 30")
		}

		if x.Age > 30 || x.Salary > 50000 {
			fmt.Println(x.Name, "makes more than 5000 or  is over than 30")
		} else {
			fmt.Println(x.Name, "makes less than 5000 and is under than 30")
		}

		if (x.Age > 30 || x.Salary < 5000) && x.FullTime {
			fmt.Println(x.Name, "matches our unclear criteria! Because his age is", x.Age, "older and his salary is= ", x.Salary)
		}
	}

}
