package main

import (
	"github.com/bcongdon/fn"
	"github.com/fogleman/gg"
)

var namer = fn.New()

func main() {
	world := NewWorld()

	nodes := make([]*Node, 10)
	for i := 0; i < 10; i++ {
		nodes[i] = NewNode(500, float64(i)*25+500)
	}
	path := NewPath(nodes)
	path.UseBrownianMotion = true

	world.paths = append(world.paths, path)

	for i := 0; i < 1000; i++ {
		world.Iterate()
	}

	dc := gg.NewContext(1000, 1000)
	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	world.Draw(dc)
	dc.Stroke()

	dc.SavePNG(namer.NameWithFileType("png"))
}
