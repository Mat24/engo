package main

import (
	"fmt"
	"learning/pkg/basic-types/functions"
	"learning/pkg/basic-types/logger"
)

const (
	const1 = "xxx"
	//Const2 _
	Const2 = 1
)

var (
	err1 = fmt.Errorf("some error")
)

//NameFunc xxx
func NameFunc(nArg1, nArg2 int, nArg3 string) (int, error) {
	acc := 0
	for i := 0; i < 10; i++ {
		acc = i + acc
	}

	return acc, nil
}

func printInit() {
	fmt.Println("Init")
}

func printEnd(i int) {
	fmt.Printf("End %d\n", i)
}

func main() {
	logger := logger.NewMyLoggerPrivate()
	printInit()
	defer printEnd(0)
	//var i int
	// bool, int, uint, int32, int64, float32, float64, complex128

	// i := int32(10) //??
	// fmt.Println(i)

	// // strings.Run()
	// logger.data = []string{}
	// basic_types.Run(logger)

	//errors.Run(logger)
	//arrays_slice.Run(logger)
	functions.Run(logger)
}
