package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Visitor interface {
	VisitDoubleExpression(*DoubleExpression)
	VisitAdditionExpression(*AdditionExpression)
}

type Expression interface {
	Visit(Visitor)
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left Expression
	right Expression
}

func (d *DoubleExpression) Visit(v Visitor) {
	v.VisitDoubleExpression(d)
}

func (a *AdditionExpression) Visit(v Visitor) {
	v.VisitAdditionExpression(a)
}

type PrintVisitor struct {
	sb strings.Builder
}

func (p *PrintVisitor) VisitAdditionExpression(a *AdditionExpression) {
	p.sb.WriteString("(")
	a.left.Visit(p)
	p.sb.WriteString("+")
	a.right.Visit(p)
	p.sb.WriteString(")")
}

func (p *PrintVisitor) VisitDoubleExpression(d *DoubleExpression) {
	p.sb.WriteString(strconv.Itoa(int(d.value)))
}

func (p *PrintVisitor) String() string {
	return p.sb.String()
}

type EvalVisitor struct {
	result float64
}

func (p *EvalVisitor) VisitAdditionExpression(a *AdditionExpression) {
	a.left.Visit(p)
	x := p.result
	
	a.right.Visit(p)
	y := p.result

	p.result = x + y
}

func (p *EvalVisitor) VisitDoubleExpression(d *DoubleExpression) {
	p.result = d.value
}

func (p *EvalVisitor) String() string {
	return strconv.Itoa(int(p.result))
}

func Visitor_Main() {
	expr := &AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			left: &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	printVisitor := &PrintVisitor{
		sb: strings.Builder{},
	}

	fmt.Printf("Visited expression: ")
	expr.Visit(printVisitor)
	fmt.Println(printVisitor)

	evalVisitor := &EvalVisitor{
		result: 0,
	}

	fmt.Printf("Eval expression: ")
	expr.Visit(evalVisitor)
	fmt.Println(evalVisitor.result)
}