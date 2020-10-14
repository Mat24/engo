package basic_types

type Logger interface {
	Log(a ...interface{}) (n int, err error)
}

type Engine struct {
	temperature float32
}

type Car struct {
	engine  Engine
	carType CarTypes
	Engine
}

type CarTypes int

func (c CarTypes) String() string {
	if c == 0 {
		return "coupe"
	}

	return "sport"
}

func Run(logger Logger) {
	car := Car{}
	car.engine.temperature = 3.14
	logger.Log(car)
	car.temperature = 1.11
	logger.Log(car)
	car.Engine.temperature = 1.455
	logger.Log(car)
	logger.Log(car.carType)
}
