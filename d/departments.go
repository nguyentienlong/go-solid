package main

import "fmt"

type Worker struct {
	ID int
	Name string
}

type Supervisor struct {
	ID int
	Name string
}

type Employee interface {
	GetID() int
	GetName() string
	Tell()
}

func(w Worker) Tell () {
	fmt.Println("im a worker")
}

func(s Supervisor) Tell () {
	fmt.Println("im a suppervisor")
}

func (w Worker) GetID() int {
	return w.ID
}

func (w Worker) GetName() string {
	return w.Name
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) GetName() string {
	return s.Name
}

type Department struct {
	Employees []Employee
}

func (d *Department) AddEmployee(e Employee) {
	d.Employees = append(d.Employees, e)
}

func (d *Department) GetEmployeeNames() (res []string) {
	for _, e := range d.Employees {
		res = append(res, e.GetName())
	}

	return res
}

func (d *Department) GetEmployeeByID(id int) Employee {
	for _, e := range d.Employees {
		if e.GetID() == id {
			return e
		}
	}

	return nil
}

func main () {

	dep := &Department{}
	dep.AddEmployee(Worker{ID:1, Name: "John"})
	dep.AddEmployee(Worker{ID:2, Name: "Doe"})

	fmt.Println(dep.GetEmployeeNames())


	e := dep.GetEmployeeByID(1)
	if e != nil {
		e.Tell()
	} else {
		fmt.Printf("not found any employee with id %d\n", 1)
	}

	e = dep.GetEmployeeByID(5)
	if e != nil {
		e.Tell()
	} else {
		fmt.Printf("not found any employee with id %d\n", 5)
	}

	//switch v:= e.(type) {
	//case Worker:
	//	fmt.Printf("found worker %+v\n", v)
	//case Supervisor:
	//	fmt.Printf("found worker %+v\n", v)
	//default:
	//	fmt.Printf("could not find an employee by id %d\n", 1)
	//}


}
