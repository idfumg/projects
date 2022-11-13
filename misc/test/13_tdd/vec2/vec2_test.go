package vec2

import "testing"

func Test_Vec2Magnitude(t *testing.T) {
	v := New(4, 0)
	if v.Magnitude() != 4 {
		t.Fail()
	}
}

func Test_Vec2Magnitude_BothAxes(t *testing.T) {
	v := New(4, -3)
	if v.Magnitude() != 5 {
		t.Fail()
	}
}
