// channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	/*
	// This will block
	<-ch
	fmt.Println("Here")
	
	go func() {
		ch <- 5
	}()
	val:= <-ch
	fmt.Println(val)*/

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Snding value ...")
			ch <-i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i:=range ch {
		val:= i
		fmt.Println("Got value ",val)
	}
}
