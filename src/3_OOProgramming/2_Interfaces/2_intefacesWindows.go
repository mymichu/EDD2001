package main

import (
	"EDD/src/3_OOProgramming/2_Interfaces/gpio"
	"EDD/src/3_OOProgramming/2_Interfaces/halWindows"
	"fmt"
)

func main() {
	halLayer := halWindows.NewWindowsLayer()
	gpio := gpio.GPIO{halLayer}
	gpio.SetPowerLED(true)
	fmt.Println(gpio.GetPowerLED())
}
