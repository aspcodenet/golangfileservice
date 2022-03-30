package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Id   int
	Age  int
	Namn string
}

type EmployeeResult struct {
	Id     int
	Salary int
}

func calculateSalary(emp Employee, ch chan EmployeeResult) {
	//Enorm beräkning
	fmt.Printf("Calculating %s\n", emp.Namn)
	time.Sleep(time.Duration(emp.Age) * time.Second)
	res := EmployeeResult{Id: emp.Id, Salary: emp.Age * 1000}
	ch <- res
}

func main() {
	allaEmployees := []Employee{}
	allaEmployees = append(allaEmployees, Employee{Id: 1, Namn: "Stefan", Age: 25})
	allaEmployees = append(allaEmployees, Employee{Id: 2, Namn: "Oliver", Age: 5})
	allaEmployees = append(allaEmployees, Employee{Id: 3, Namn: "Josefine", Age: 12})

	// CHANNEL kanal för kommunikation
	ch := make(chan EmployeeResult)

	for _, employee := range allaEmployees {
		go calculateSalary(employee, ch)
		//fmt.Printf("Lönen för %d blev %d", employee.Id, salary)
	}
	fmt.Println("Nu är alla goroutines igång...jag väntar på svar nu")
	for {
		employeeResult := <-ch
		fmt.Printf("Lönen för %d blev %d", employeeResult.Id, employeeResult.Salary)
	}

}
