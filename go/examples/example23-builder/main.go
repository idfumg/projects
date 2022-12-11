package main

import (
	"fmt"
)

const (
	LEFT  = 1 << 0
	RIGHT = 1 << 1
	TOP   = 1 << 2
	BACK  = 1 << 3
)

type Shield struct {
	left  bool
	top   bool
	right bool
	back  bool
}

type ShieldBuilder struct {
	code int
}

func NewShieldBuilder() *ShieldBuilder {
	return new(ShieldBuilder)
}

func (sb *ShieldBuilder) RaiseLeft() *ShieldBuilder {
	sb.code |= LEFT
	return sb
}

func (sb *ShieldBuilder) RaiseRight() *ShieldBuilder {
	sb.code |= RIGHT
	return sb
}

func (sb *ShieldBuilder) RaiseTop() *ShieldBuilder {
	sb.code |= TOP
	return sb
}

func (sb *ShieldBuilder) RaiseBack() *ShieldBuilder {
	sb.code |= BACK
	return sb
}

func (sb *ShieldBuilder) Build() *Shield {
	return &Shield{
		left:  sb.code&LEFT > 0,
		right: sb.code&RIGHT > 0,
		top:   sb.code&TOP > 0,
		back:  sb.code&BACK > 0,
	}
}

func main() {
	builder := NewShieldBuilder()
	shield := builder.RaiseLeft().RaiseTop().Build()
	fmt.Printf("%+v\n", shield)
}
