package gpio

import (
	"EDD/src/3_OOProgramming/2_Interfaces/halAPI"
)

type GPIO struct {
	HalAPI halAPI.HalApi
}

func (m *GPIO) SetPowerLED(on bool) {
	m.HalAPI.SetRawGPIO(20, true)
}

func (m *GPIO) GetPowerLED() (bool, error) {
	return m.HalAPI.GetRawGPIO(20)
}
