// struct demo
package main

import (
	"fmt"
)

type Person struct {
	Name 	string  // name of person
	Gender  string  // male or female
	Age 	int     // age of person
}

func main() {
	p1 := Person{"John", "male", 21}
	fmt.Println(p1)

	fmt.Printf("%+v\n",p1)
	fmt.Println(p1.Name)

	p2 := Person{
		Name: "Emma",
		Age: 20,
		Gender:  "female",
	}
	fmt.Printf("%+v\n",p2)

	p3:=Person{}
	fmt.Printf("%+v\n",p3)
}
