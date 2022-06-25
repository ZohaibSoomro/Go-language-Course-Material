// Method demo
package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name 	string  // name of person
	Gender  string  // male or female
	Age 	int     // age of person
}

// NewPerson will create a new person and will validate the input
func NewPerson(name string, gender string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name can't be empty")
	}

	if gender == "" {
		return nil, fmt.Errorf("gender can't be empty")
	}

	if age <= 0 {
		return nil, fmt.Errorf("age must be a positive integer")
	}

	person := &Person{
		Name: name,
		Gender: gender,
		Age:  age,
	}
	return person, nil
}


//IsAdult() tells if a person is an adult 
func (p *Person) IsAdult() bool{
	if p.Age<18{
		return false
	}
	return true	
}

func main() {
	p, err := NewPerson("John", "male", 21)

	if err != nil {
		fmt.Printf("error: can't create person object - %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n",p)
	fmt.Println("Adult: ",p.IsAdult())
}
