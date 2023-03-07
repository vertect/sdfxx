package primitive

import (
	"math"

	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
)

func RegularPolygon2D(vertices int, radius float64) (sdf.SDF2, error) {
	vertex := make([]v2.Vec, vertices)
	theta := math.Pi * 2 / float64(vertices)
	for i := 0; i < vertices; i++ {
		t := theta * float64(i)
		vertex[i] = v2.Vec{
			X: radius * math.Cos(t),
			Y: radius * math.Sin(t),
		}
	}
	return sdf.Polygon2D(vertex)
}
