package basic_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type loggerMock struct {
	counter int
}

func (l *loggerMock) Log(a ...interface{}) (n int, err error) {
	l.counter++
	return 0, err
}

func TestMock(t *testing.T) {
	//loggerMock := &loggerMock{}
	//Run(loggerMock)
	//assert.Equal(t, 4, loggerMock.counter)
}

func TestRun(t *testing.T) {
	car := Car{}
	//Esto no es así
	assert.IsType(t, Engine{}, car)
}
