package bridge

import "fmt"

//1.- Interface render have common function
type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	//implement what a vector have
}

//2.- This struct must have same function than renderer
func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

//3.- This struct must have same function than renderer
func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

//4.- Now we just send Renderer interface and now we can pass to this function or VectorRender or RasterRender struct
func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}
