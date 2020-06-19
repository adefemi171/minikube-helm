package main

import "fmt"

func main(){
	count := make(chan int)

	go func(){
		current := 0
		for {
			current = <-count
			current--
			count <- current
			fmt.Println(count)
		}
	}()
	count <- 2
	fmt.Println("+", <-count)
}