package main

import (
	"fmt"
	"time"
	"math"
	"github.com/trosh/term"
)

func Triangle(center term.Pxl,
              radius float64,
              angle  float64) (term.Pxl, term.Pxl, term.Pxl) {
	angle2 := angle + 2 * math.Pi / 3
	angle3 := angle + 4 * math.Pi / 3
	p1 := term.Pxl{
		int(math.Sin(angle)*radius * 1.7) + center.X,
		int(math.Cos(angle)*radius      ) + center.Y}
	p2 := term.Pxl{
		int(math.Sin(angle2)*radius * 1.7) + center.X,
		int(math.Cos(angle2)*radius      ) + center.Y}
	p3 := term.Pxl{
		int(math.Sin(angle3)*radius * 1.7) + center.X,
		int(math.Cos(angle3)*radius) + center.Y}
	return p1, p2, p3
}

func main() {
	s := term.Scr{term.Pxl{1, 1}, term.Pxl{60, 28}, 0}
	s.Flush()
	s = term.Scr{term.Pxl{10, 7}, term.Pxl{40, 22}, 52}
	orig := term.Pxl{25, 15}
	radius := 13.0
	for i := 0.0; ; i += 0.1 {
		p1, p2, p3 := Triangle(orig, radius, math.Sin(i) + i)
		s.Line(p1, p2, true)
		s.Line(p2, p3, true)
		s.Line(p3, p1, true)
		fmt.Printf("\033[0;0H")
		time.Sleep(25 * time.Millisecond)
		s.Flush()
	}
	fmt.Printf("\033[0m\n")
}
