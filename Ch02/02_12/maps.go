// Go's map data structure
package main

import (
	"fmt"
)

func main() {
	heights := map[string]float64{
		"Ahmed": 5.2,
		"Ali": 6.1,
		"Aslam": 5.5, // Must have trailing comma in multi line
	}
	// Number of items
	fmt.Println(heights)

	//to know the length of a map
	fmt.Println(len(heights))

	// Get a value
	fmt.Println(heights["Ali"])	

	// Get zero value if not found
	fmt.Println(heights["Akbar"]) // 0

	// Use two value to see if found
	value, ok := heights["Akbar"]
	if !ok {
		fmt.Println("Akbar not found")
	} else {
		fmt.Println(value)
	}

	// Set 
	heights["Akbar"] = 6.3
	fmt.Println(heights)

	// Delete
	delete(heights, "Ali")
	fmt.Println(heights)

	// Single value "for" is on keys
	for key := range heights {
		fmt.Println(key)
	}

	// Double value "for" is on keys,values
	for key,value := range heights {
		fmt.Println(key,"->",value)
	}
}
