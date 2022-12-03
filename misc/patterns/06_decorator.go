package main

import "fmt"

type Shape6 interface {
	Render() string
}

type Circle6 struct {
	Radius float32
}

func (c *Circle6) Render() string {
	return fmt.Sprintf("Circle of radius: %f", c.Radius)
}

func (c *Circle6) Resize(factor float32) {
	c.Radius *= factor
}

type Square6 struct {
	Side float32
}

func (s *Square6) Render() string {
	return fmt.Sprintf("Square with a side: %f", s.Side)
}

type ColoredShape struct {
	Shape Shape6
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s: has the color %s", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape        Shape6
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %.2f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func Decorator() {
	circle := Circle6{2}
	fmt.Println(circle.Render())
	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())
	transparentCircle := TransparentShape{&redCircle, 0.9}
	fmt.Println(transparentCircle.Render())
}
