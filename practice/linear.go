// Linear Search

package main

import (
	"fmt"
)

func linearSearch(dataList []int,key int) bool{
   
   for _,item := range dataList{
		 if item == key{
			 return true
		 }
	 }
	 return false
}


func main5(){
    
	var dataList []int = []int{1,2,3,453,2,34,2,8}
   fmt.Println(linearSearch(dataList,8))
}