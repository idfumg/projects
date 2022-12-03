package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name     string
	Color    string
	Children []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')
	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle5(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare5(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

func Composite() {
	drawing := GraphicObject{"My drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewCircle5("Red"))
	drawing.Children = append(drawing.Children, *NewSquare5("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle5("Blue"))
	group.Children = append(group.Children, *NewSquare5("Blue"))

	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
}
