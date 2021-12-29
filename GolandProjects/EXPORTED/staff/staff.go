package staff

import "log"

var OverPaidLimit = 75000
var UnderPaidLimit = 50000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (employee *Office) All() []Employee {
	return employee.AllStaff
}

func (employee *Office) Overpaid() []Employee {
	var overPaid []Employee
	for _, x := range employee.AllStaff {
		if x.Salary > OverPaidLimit {
			overPaid = append(overPaid, x)
		}
	}
	return overPaid
}

func (employee *Office) UnderPaid() []Employee {
	var underpaid []Employee
	for _, x := range employee.AllStaff {
		if x.Salary < UnderPaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}

func (employee *Office) notVisible() {
	log.Println("Hello, Mundo!")
}
