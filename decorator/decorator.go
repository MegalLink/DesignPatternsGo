package decorator

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

// THIS RESIZE METHOD IS JUST FROM CIRCLE SO WE CAN'T ADD THIS METHOD TO Shape Interface
func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// we have this structs above this their methods but what if we want this shapes to have a color
// we can add color property to circle and sauare but we dont know how is this going to affect the behavior
// so we are violating solid principles

// posible solution implement aggregate, the problem is we have to create new struct for every shape we could have
/*
type ColoredSquare struct {
	Square
	Color string
}
*/

// solution here is to have decorator
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

// other decorator
type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has transparency %d percent", t.Shape.Render(), int(t.Transparency*100.0))
}
