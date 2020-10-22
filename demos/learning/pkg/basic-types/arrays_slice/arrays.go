package arrays_slice

import "fmt"

type Logger interface {
	Log(a ...interface{}) (n int, err error)
}

func fooSlice(sInt []int) {
	sInt[0] = 1000
}

func fooArrayPtr(sInt [3]*int) {
	i := 800
	sInt[0] = &i
}

func fooSlicePtr(sInt []*int) {
	i := 888
	sInt[0] = &i
}

func fooArray(sInt [3]int) {
	sInt[0] = 1000
}

func fooInt(i int) {
	i = 999
}

func printSlicePtr(s []*int, logger Logger) {
	for _, v := range s {
		logger.Log(*v)
	}
}

func runArrays(logger Logger) {
	{

		arr := [3]int{}

		slice := make([]int, 6, 10)

		logger.Log("slice", slice, "len", len(slice), "cap", cap(slice))
		arr[0] = 9
		slice[0] = 8
		slice2 := make([]int, 3)
		slice2[1] = 2
		slice2[2] = 3

		count := cap(slice)
		logger.Log("cap", count)

		for i, item := range slice2 {
			slice[3+i] = item
			logger.Log("slice", slice, "len", len(slice), "cap", cap(slice))
		}

		slice = append(slice, make([]int, 6)...)
		count = cap(slice)
		logger.Log("cap", count)

		logger.Log(arr)
		fooArray(arr)
		logger.Log(arr)
		logger.Log("******")
		logger.Log(slice)
		fooSlice(slice)
		fooInt(slice[0])
		logger.Log(slice)
		logger.Log("//////")

		arrayPtr := [3]*int{}
		logger.Log(arrayPtr)
		fooArrayPtr(arrayPtr)
		logger.Log(arrayPtr)
		p := 8
		slicePtr := []*int{&p}

		logger.Log("------")
		printSlicePtr(slicePtr, logger)
		fooSlicePtr(slicePtr)
		printSlicePtr(slicePtr, logger)
	}

	org := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//slice := make([]int, 0, 10)
	slice := make([]int, 0, 10)
	logger.Log(org)
	logger.Log(slice)

	for _, v := range org {
		slice = append(slice, v)
	}

	logger.Log(slice)

}

type myStruct struct {
	str string
}

func runMaps(logger Logger) {
	m := map[myStruct]int{
		{str: "key"}:  1,
		{str: "key2"}: 2,
	}
	m[myStruct{str: "key3"}] = 3
	logger.Log(m)

	if i, ok := m[myStruct{str: "x"}]; ok {
		logger.Log(i)
	} else {
		logger.Log("not found")
	}

	for k, v := range m {
		logger.Log(fmt.Sprintf("k:%v,v:%v", k, v))
	}

}

func Run(logger Logger) {
	runMaps(logger)
	runArrays(logger)
}
