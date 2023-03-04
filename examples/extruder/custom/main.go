package main

import (
	"fmt"
	"path"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	"github.com/vertect/sdfxx/extruder"
)

func main() {
	factories := map[string]extruder.ExtrudeFuncFactory{
		"cos":   extruder.CosExtrude,
		"circ":  extruder.CircExtrude,
		"quad":  extruder.QuadExtrude,
		"cubic": extruder.CubicExtrude,
		"quart": extruder.QuartExtrude,
		"quint": extruder.QuintExtrude,
		"expo":  extruder.ExpoExtrude,
	}
	box := sdf.Box2D(v2.Vec{X: 10, Y: 10}, 0)
	for name, factory := range factories {
		model := extruder.CustomScaleExtrude3D(box, 20, v2.Vec{X: 2, Y: 3}, factory)
		render.ToSTL(
			model,
			path.Join("./out", fmt.Sprintf("extruder_custom_%s.stl", name)),
			render.NewMarchingCubesUniform(300),
		)
	}
}
