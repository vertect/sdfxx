package main

import (
	"path"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	"github.com/vertect/sdfxx/primitive"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	p, err := primitive.RegularPolygon2D(8, 10)
	if err != nil {
		return err
	}
	render.ToSTL(
		sdf.Extrude3D(p, 1),
		path.Join("./out", "primitive_regular_polygon.stl"),
		render.NewMarchingCubesUniform(300),
	)
	return nil
}
