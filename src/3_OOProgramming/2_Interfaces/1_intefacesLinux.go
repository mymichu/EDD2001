package main

import (
	"EDD/src/3_OOProgramming/2_Interfaces/gpio"
	"EDD/src/3_OOProgramming/2_Interfaces/halLinux"
	"fmt"
)

func main() {
	halLayer := halLinux.NewLinuxLayer()
	gpio := gpio.GPIO{halLayer}
	gpio.SetPowerLED(true)
	fmt.Println(gpio.GetPowerLED())
}
