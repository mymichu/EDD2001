package halAPI

type HalApi interface {
	SetRawGPIO(number uint8, value bool) error
	GetRawGPIO(number uint8) (bool, error)
}
