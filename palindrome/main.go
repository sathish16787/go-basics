package main

import (
	"fmt"
	"os"
)

func main() {
	var num, reserved int
	fmt.Println("Enter a number to check panlindrome: ")

	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if num <= 0 {
		fmt.Println("Enter a valid Integer, Greater than Zero")
	}

	original := num

	for num > 0 {
		last_digit := num % 10
		reserved = reserved*10 + last_digit
		num = num / 10
	}

	if original == reserved {
		fmt.Printf("The given number is %d  is  a palindrome\n", original)
	} else {
		fmt.Printf("The given number is %d  is not a palindrome\n", original)
	}
}
