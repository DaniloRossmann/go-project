package operations

import (
	"errors"
)

//Diff é um função que subtrai dois valores passados por paranmetro
func Diff(a int, b int) (int, error) {
	res := a - b
	if res < 0 {
		return 0, errors.New("Número negativo")
	} else {
		return res, nil
	}

}
