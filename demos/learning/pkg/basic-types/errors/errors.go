package errors

import (
	"errors"
	"fmt"
)

var (
	errZeroByZeroDiv = errors.New("unable to div 0 by 0")
)

type Logger interface {
	Log(a ...interface{}) (n int, err error)
}

func div(a, b float32) (float32, error) {
	if a == 0 && b == 0 {
		return 0, errZeroByZeroDiv
	}

	return a / b, nil
}

func div2(a, b float32, r *float32) error {
	if a == 0 && b == 0 {
		return errZeroByZeroDiv
	}

	*r = a / b
	return nil
}

func Run(logger Logger) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nRecovered in f\n", r)
		}
	}()

	// r, err := div(0, 0)
	// if err != nil {
	// 	logger.Log(err.Error())
	// 	return
	// }

	// logger.Log(fmt.Sprintf("res %f", r))

	s := []float32{1, 2, 3, 0}
	p := float32(0)
	var err error

	for _, value := range s {
		if err = div2(0, value, &p); err != nil {
			logger.Log(err.Error())
		}
	}

}
