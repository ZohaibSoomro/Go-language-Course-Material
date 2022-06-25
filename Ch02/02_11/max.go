// Calculate maximal value in a slice
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15,67}
	max:=nums[0]

	for _,value:= range nums[1:]{
		if value >max{
			max=value
		}
	}

	fmt.Println(max)
}
