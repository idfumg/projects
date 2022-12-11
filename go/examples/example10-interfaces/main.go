package main

import (
	"errors"
	"fmt"
)

type Walker interface {
	walk(p Point) error
	getPosition() Point
}

type Talker interface {
	talk(s string)
}

type Point struct {
	x int
	y int
}

type Human struct {
	Position Point
}

type Animal struct {
	Position Point
}

func (h *Human) walk(p Point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("Wrong point")
	}
	h.Position = p
	fmt.Println("Human walked to", p)
	return nil
}

func (h *Human) getPosition() Point {
	return h.Position
}

func (h *Human) talk(s string) {
	fmt.Println("Human talking", s)
}

func (h *Animal) walk(p Point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("Wrong point")
	}
	h.Position = p
	fmt.Println("Animal walked to", p)
	return nil
}

func (h *Animal) getPosition() Point {
	return h.Position
}

func (h *Animal) talk(s string) {
	fmt.Println("Animal talking", s)
}

func move(w Walker, points []Point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func moveHuman(h Human, points []Point) error {
	for _, p := range points {
		err := h.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	steps := []Point{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	jack := Human{}
	err := move(&jack, steps)
	if err != nil {
		fmt.Println(err)
	}

	dog := Animal{}
	err = move(&dog, steps)
	if err != nil {
		fmt.Println(err)
	}

	moveHuman(jack, steps)
	// moveHuman(dog, steps)
}
