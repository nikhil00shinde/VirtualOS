package  main


import (
	"fmt"
)

type Car struct{
	Model string
	Mileage int
	Brand string
}

func main(){

var car = Car{"X5",43,"BMW"}
fmt.Println(car.Model)

 car.Model = "M3"
fmt.Println(car)
}