package main

import (
	"EDD/src/3_OOProgramming/1_Struct/Car"
	"fmt"
)

func main() {
	var car Car.Car
	car.SetColor("yellow")
	car.SetPosition(0, 0)
	car.SetPosition(10, 20)
	car.SetPosition(15, 30)
	fmt.Println(car.GetColor())
	fmt.Println(car.GetPositionLogEntry(-1))
}
