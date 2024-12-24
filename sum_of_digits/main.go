package main

import (
	"fmt"
	"os"
)

func main()  {
	var number,sum int
	fmt.Println("Enter the number: ")

	_,err:= fmt.Scanf("%d",&number)
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	if number <= 0 {
		fmt.Println("Enter valid positive integer")
		os.Exit(1)
	}

	original := number
	for number>0{
		last_digit := number%10
		sum = sum + last_digit
		number = number/10
	}
 fmt.Printf("sum of all digits %d is %d \n",original,sum)
}