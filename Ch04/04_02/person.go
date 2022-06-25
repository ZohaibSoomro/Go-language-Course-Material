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

//GetAge() returns age of person
func (p *Person) IsAdult() bool{
	if p.Age<18{
		return false
	}
	return true	
}

func main() {
	p1 := Person{"John", "male", 17}
	p2 := Person{
		Name: "Emma",
		Age: 20,
		Gender:  "female",
	}
	fmt.Printf("Person1: %+v\n",p1)
	fmt.Println("Adult: ",p1.IsAdult())

	fmt.Printf("Person2: %+v\n",p2)
	fmt.Println("Adult: ",p2.IsAdult())
}
