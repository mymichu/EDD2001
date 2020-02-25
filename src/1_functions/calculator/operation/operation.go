package operation

import "errors"

func Division(numerator int16, denominator int16) float32 {
	result := float32(numerator) / float32(denominator)
	return result
}

func DivisionWithNullCheck(numerator int16, denominator int16) (float32, error) {
	var err error = nil
	var result float32 = 0.0

	if denominator != 0 {
		result = float32(numerator) / float32(denominator)
	} else {
		err = errors.New("denominator is value 0")
	}

	return result, err
}
