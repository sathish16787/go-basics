package main

import (
	"fmt"
	"os"
)

func main() {

	var num, reversed int

	fmt.Println("Enter a number to reverse it: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	original := num
	for num > 0 {

		last_digit := num % 10
		reversed = reversed*10 + last_digit
		num = num / 10

	}
	fmt.Printf("The reverse of %d is %d\n", original, reversed)
}
