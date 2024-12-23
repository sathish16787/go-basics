package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	sum := 1
	fmt.Println("Enter a positive integer")

	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if a < 1 {
		fmt.Println("Enter valid input....Enter greater than zero")
	} else {
		for i := 2; i <= a; i++ {
			sum = sum * i
		}
		fmt.Println(sum)
	}

}
