package main

import "fmt"

func main() {

	var a, b int
	fmt.Println("Enter the value of a and b : ")
	fmt.Scanf("%d %d", &a, &b)

	if a > b {
		fmt.Printf("%d is maximum of two numbers \n", a)
	} else if b > a {
		fmt.Printf("%d is maximum of two numbers \n", b)
	} else {
		fmt.Println("Both the numbers are equal\n")
	}

}
