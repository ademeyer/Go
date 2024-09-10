package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary float64
}

func createPerson(name string, age int, job string, salary float64) *Person {
	return &Person{name, age, job, salary}
}

func (p *Person) displayInfo() {
	fmt.Printf("Name: %s\nAge: %d\nJob: %s\nSalary: %.2f\n\n", p.name, p.age, p.job, p.salary)
}

func main() {
	var people [10]Person

	for i := 0; i < len(people); i++ {
		people[i] = *createPerson(fmt.Sprintf("Person %d", i+1), i+20, fmt.Sprintf("Job %d", i+1), 5000.00+float64(i+1)*1000.00)
	}

	for i := 0; i < len(people); i++ {
		people[i].displayInfo()
	}
}
