package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	City string `json:"city"`
}

type Salary struct {
	Basic, HRA, TA float64
}
type EmployeeDetails struct {
	FirstName, LastName, Email string
	Age int
	MonthlySalary []Salary
}

func (e EmployeeDetails) EmplDetlInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Email)
	fmt.Println(e.Age)
	for i, info := range e.MonthlySalary {
		fmt.Println("====================")
		fmt.Println(i, info.Basic, info.HRA, info.TA)
	}
	return "--------------------"
}

func main() {
	var employeeA = new(Employee)
	var jsonas string = `{
"firstname":"Bob",
"lastname": "NX",
"city": "Xi'an"
}`
	json.Unmarshal([]byte(jsonas), employeeA)
	fmt.Println(employeeA)

	var employeeB = new(Employee)
	employeeB.FirstName = "Alice"
	employeeB.LastName = "Bitch"
	employeeB.City = "XiaMen"

	var jsonbs, _ = json.Marshal(employeeB)
	fmt.Println(jsonbs)
	fmt.Printf("%s\n", employeeB)

	e := EmployeeDetails{
		FirstName:"Jack",
		LastName:"Jones",
		Email:"jack-jones@xxx.xxx",
		Age: 28,
		MonthlySalary:[]Salary{
			Salary{
				Basic:1500.00,
				HRA:500.00,
				TA:200.00,
			},
			Salary{
				Basic:1500.00,
				HRA:600.00,
				TA:200.00,
			},
		},
	}
	fmt.Println(e)
	e.MonthlySalary = append(e.MonthlySalary, Salary{Basic:1000.00,HRA:300.00,TA:300.00})
	fmt.Println(e)
	fmt.Println(e.EmplDetlInfo())
}
