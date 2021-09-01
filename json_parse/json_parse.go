package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for _, user := range users.Users {
		fmt.Println("User Type: " + user.Type)
		fmt.Println("User Age: " + strconv.Itoa(user.Age))
		fmt.Println("User Name: " + user.Name)
		fmt.Println("Facebook Url: " + user.Social.Facebook)
	}

	jsonFileEmp, err := os.Open("employees.json")
	if err != nil {
		log.Fatal(err)
	}

	byteValueEmp, _ := ioutil.ReadAll(jsonFileEmp)

	var emps Employees

	json.Unmarshal(byteValueEmp, &emps)

	fmt.Println(len(emps.Employees))

	for i := 0; i < len(emps.Employees); i++ {
		fmt.Println("Name: " + emps.Employees[i].Name)
	}

	// for _, emp := range emps.Employees {
	// 	fmt.Println(emp)
	// }
	// if err := json.NewDecoder(jsonFile).Decode(&emps); err != nil {
	// 	for _, emp := range emps.Employees {
	// 		fmt.Println(emp)
	// 	}
	// }
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}

type Employee struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"Age"`
}
