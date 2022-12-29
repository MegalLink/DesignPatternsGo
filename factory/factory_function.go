package factory

import (
	"fmt"

	"github.com/MegalLink/design-patterns/logger"
)

// factory interface expose just the interface methods
type IPersonService interface {
	SayHello()
}

type Person struct {
	name    string
	age     int
	isOlder bool
}

// factory function is used to define default values.
func NewDefaultPerson(name string, age int, logger logger.IFastLogger) IPersonService {
	isOlder := false
	if age > 18 {
		isOlder = true
	}
	fmt.Println(logger)

	return &Person{
		name:    name,
		age:     age,
		isOlder: isOlder,
	}
}

func (p *Person) SayHello() {
	fmt.Println(p)
}
