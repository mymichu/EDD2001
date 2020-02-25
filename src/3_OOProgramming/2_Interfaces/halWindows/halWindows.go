package halWindows

import (
	"EDD/src/3_OOProgramming/2_Interfaces/halAPI"
	"fmt"
)

type linuxLayer struct {
	gpioMap map[uint8]bool
}

func (r *linuxLayer) SetRawGPIO(number uint8, on bool) error {
	fmt.Println("WRITE WINDOWS Nr: ", number, " VALUE: ", on)
	r.gpioMap[number] = on
	return nil
}

func (r linuxLayer) GetRawGPIO(number uint8) (bool, error) {
	fmt.Println("READ WINDOWS Nr: ", number)
	return r.gpioMap[number], nil
}

func NewWindowsLayer() halAPI.HalApi {
	return &linuxLayer{gpioMap: make(map[uint8]bool)}
}
