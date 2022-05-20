// Bar Chart

package main

import (
	"fmt"
)

func printBar(arr [5] int){
	var max int = 0
	for i:=0;i<len(arr);i++{
		if max<arr[i]{
			max = arr[i]
		}
	}

	for i:=max;i>0;i--{
      for j:=0;j<len(arr);j++{
				if arr[j]>= i{
					fmt.Print("*\t")
				}else{
					fmt.Print("\t")
				}
			}
			fmt.Println()
	}
}

func main2(){
	var arr [5]int = [5]int{12,34,4,4,5}
	printBar(arr)
	printBar([5]int{1,2,3,4,5})

}