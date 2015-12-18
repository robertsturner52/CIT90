package main

import (
	"fmt"
)

func main() {
	weeksPerYear := 40
	hoursPerDay := 8
	salaryPerHour := 18.75
	annualSalary := float64(weeksPerYear) * float64(hoursPerDay) * salaryPerHour
	fmt.Println("His salary was $", annualSalary, "a year.")
}
