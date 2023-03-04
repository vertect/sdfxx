package main

import (
	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
	"github.com/vertect/sdfxx/extruder"
)

func main() {
	box := sdf.Box2D(v2.Vec{X: 10, Y: 10}, 0)
	model := extruder.SlideExtrude3D(box, v3.Vec{X: 20, Y: 0, Z: 20})
	render.ToSTL(
		model,
		"./out/extruder_slide.stl",
		render.NewMarchingCubesUniform(300),
	)
}
