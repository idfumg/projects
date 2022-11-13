package vec2

import "math"

type Vec2 struct {
	x float64
	y float64
}

func New(x, y float64) *Vec2 {
	return &Vec2{
		x: x,
		y: y,
	}
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
