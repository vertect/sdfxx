package extruder

import (
	"math"

	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

type ExtrudeSDF3 struct {
	sdf     sdf.SDF2
	height  float64
	extrude sdf.ExtrudeFunc
	bb      sdf.Box3
}

// Evaluate returns the minimum distance to an extrusion.
func (s *ExtrudeSDF3) Evaluate(p v3.Vec) float64 {
	// sdf for the projected 2d surface
	a := s.sdf.Evaluate(s.extrude(p))
	// sdf for the extrusion region: z = [-height, height]
	b := math.Abs(p.Z) - s.height
	// return the intersection
	return math.Max(a, b)
}

// SetExtrude sets the extrusion control function.
func (s *ExtrudeSDF3) SetExtrude(extrude sdf.ExtrudeFunc) {
	s.extrude = extrude
}

// BoundingBox returns the bounding box for an extrusion.
func (s *ExtrudeSDF3) BoundingBox() sdf.Box3 {
	return s.bb
}
