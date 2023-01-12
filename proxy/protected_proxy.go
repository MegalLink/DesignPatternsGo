package proxy

import "fmt"

// to avoid chaingin code
// provide same interface but with different behavior

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

// we want to call Drive only when car has a driver so we can build a protecction proxy

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		car: Car{}, driver: driver,
	}
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		fmt.Println("Driver yo young")
	}
}
