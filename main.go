package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	emps := []User{
		{"John", "23"},
	}
}
