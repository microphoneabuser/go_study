package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 1
	// p := &x
	// fmt.Println(p)
	// a := *p
	// fmt.Println(a)

	// g := new(int)
	// fmt.Println(g)

	// arr := []int{1, 2, 3, 4}
	// fmt.Printf("%ds", arr)

	// a := []int{1, 2, 3, 4}
	// fmt.Println(a)
	// fmt.Println(&a[0])
	// a = append(a, 5)
	// fmt.Println(a)
	// fmt.Println(&a[0])

	a := [5]int{1, 2, 3, 4, 5}
	b := a[2:]
	b[1] = 6
	b = append(b, 7)
	fmt.Println(a, b)

	emp := Employee{ID: 1, Name: "Ivan", Address: "djskl", Position: "Software Engineer"}
	fmt.Println(emp)

	changeEmp(&emp)

	fmt.Println(emp)
}

type Employee struct {
	ID       int
	Name     string
	Address  string
	DoB      time.Time
	Position string
	Salary   int
}

func changeEmp(emp *Employee) {
	emp.Name = "loh"
}
