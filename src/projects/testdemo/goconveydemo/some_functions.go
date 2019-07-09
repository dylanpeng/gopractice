package goconveydemo

import "errors"

func IsEqual(a, b int) bool {
	return a == b
}

func IsEqualWithErr(a, b int) (bool, error) {
	if a > b {
		return false, errors.New("over")
	} else if a < b {
		return false, errors.New("under")
	} else {
		return true, nil
	}
}
