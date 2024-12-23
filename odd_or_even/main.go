package main

import (
	"fmt"
	"os"
)

func main()  {
	 var  a int

	 fmt.Println("Enter a number: ")

	 _ ,err := fmt.Scanf("%d",&a)
	 if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	 }
	 fmt.Printf("You entered %d \n",a)



   if a<=0 {
		fmt.Println("Enter number greater than Zero")
	 }else if a%2==0{
		fmt.Println("The given number is even")
	 } else {
		fmt.Println("The given number is odd")
	 }


}