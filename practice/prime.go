// Print Prime Number in a given range

package main

import (
	"fmt"
	"math"
)

func printPrimeNumber(num1 , num2 int){
	if num1 <2 || num2 <2{
		fmt.Println("Numbers must be greater than 2")
		return 
	}

	for num1<=num2{
		isPrime := true
		for i:=2 ; i <= int(math.Sqrt(float64(num1))); i++{
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
   
		if isPrime{
			fmt.Printf("%d is Prime Number\n",num1)
		}
		num1++
	}
}

func main1(){
	fmt.Println("Print Prime Number from a Given Range")
  printPrimeNumber(2,10)
}