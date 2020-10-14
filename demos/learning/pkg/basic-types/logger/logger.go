package logger

import "fmt"

type myLoggerPrivate struct {
	data    []string
	counter int
}

//type ExportedLogger = myLoggerPrivate

func NewMyLoggerPrivate() *myLoggerPrivate {
	return &myLoggerPrivate{
		data: []string{},
	}
}

func (l *myLoggerPrivate) Log(a ...interface{}) (n int, err error) {
	return fmt.Println(a)
}

// func (l *myLoggerPrivate) Log(a ...interface{}) (n int, err error) {
// 	if l.counter++; l.counter == 1 {
// 		l.flush()
// 		l.data = []string{}
// 	}
// 	l.data = append(l.data, fmt.Sprintf("%+v", a))
// 	return 0, nil
// }

// func (l *myLoggerPrivate) flush() (n int, err error) {
// 	for _, str := range l.data {
// 		fmt.Println(str)
// 	}

// 	return 0, nil
// }
