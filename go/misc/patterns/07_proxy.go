package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    *Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		fmt.Println("Driver is too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		&Car{},
		driver,
	}
}

// the client will always get some car proxy class with specific logic
func Proxy_Protection() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
}
