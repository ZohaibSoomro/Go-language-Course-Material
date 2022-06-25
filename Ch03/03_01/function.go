// Basic function definition
package main

import (
	"fmt"
)

func add(a int ,b int ) int {
	return a+b
}

func divmod(a int,b int) (int,int){
	return a/b,a%b
}

func main() {
	val:=add(2,3)
	fmt.Println(val)

	quot,rem:=divmod(7,2)
	fmt.Println("quotient ",quot)
	fmt.Println("remainder ",rem)
}
