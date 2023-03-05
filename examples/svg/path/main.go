package main

import (
	"os"

	"github.com/deadsy/sdfx/render"
	"github.com/deadsy/sdfx/sdf"
	"github.com/vertect/sdfxx/svg"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	f, err := os.OpenFile("./Ghostscript_Tiger.svg", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	svg, err := svg.New(f, &svg.Option{})
	if err != nil {
		return err
	}
	g, err := svg.Find("g58")
	if err != nil {
		return err
	}
	path, err := g.Draw()
	if err != nil {
		return err
	}
	model := sdf.Extrude3D(path, 10)
	render.ToSTL(
		model,
		"./out/tiger.stl",
		render.NewMarchingCubesUniform(300),
	)
	return nil
}
