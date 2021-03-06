package main

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Employee struct {
	Id   int
	Age  int
	Namn string
	City string
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

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Employee{})
	// for _, employee := range allaEmployees {
	// 	db.Create(&employee)
	// 	//insert into Employee
	// 	//reflection runtime ta reda på vilka fält som finns
	// }

	//Lista alla
	ret := []Employee{}
	db.Find(&ret, &Employee{})
	for _, device := range ret {
		fmt.Println(device.Namn)
	}

	ret1 := Employee{}
	db.First(&ret1, "age=?", 12)

	db.First(&ret1, 2)
	ret1.City = "Test"
	ret1.Age = 18
	db.Model(&ret1).Updates(ret1)

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
