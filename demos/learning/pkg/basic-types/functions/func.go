package functions

import "fmt"

type Logger interface {
	Log(a ...interface{}) (n int, err error)
}

type TX struct {
	id int `json="index"`
}

type Operator func(x, y int) int

func Adder(x, y int) int {
	return x + y
}

func MakeAdder(l Logger) Operator {
	return func(x, y int) int {
		l.Log(fmt.Sprintf("adding x %v y %v ", x, y))
		return x + y
	}
}

func Run(l Logger) {
	adder := MakeAdder(l)
	r := adder(1, 2)
	l.Log(fmt.Printf("result = %v", r))
}
