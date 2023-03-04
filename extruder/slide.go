package extruder

import (
	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

func SlideExtrude3D(sdf2 sdf.SDF2, translate v3.Vec) sdf.SDF3 {
	s := ExtrudeSDF3{}
	s.sdf = sdf2
	s.height = translate.Z / 2
	s.extrude = extrude(translate)

	// work out the bounding box
	bb := sdf2.BoundingBox()
	d := v2.Vec{X: translate.X, Y: translate.Y}.DivScalar(2)
	ba := bb.Translate(d.MulScalar(-1)).Extend(bb.Translate(d))
	s.bb = sdf.Box3{
		Min: v3.Vec{X: ba.Min.X, Y: ba.Min.Y, Z: -s.height},
		Max: v3.Vec{X: ba.Max.X, Y: ba.Max.Y, Z: s.height},
	}
	return &s
}

func extrude(translate v3.Vec) sdf.ExtrudeFunc {
	t := v2.Vec{X: translate.X, Y: translate.Y}.DivScalar(translate.Z)
	return func(p v3.Vec) v2.Vec {
		return v2.Vec{X: p.X, Y: p.Y}.Add(t.MulScalar(p.Z))
	}
}
