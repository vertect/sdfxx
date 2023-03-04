package extruder

import (
	"math"

	"github.com/deadsy/sdfx/sdf"
	v2 "github.com/deadsy/sdfx/vec/v2"
	v3 "github.com/deadsy/sdfx/vec/v3"
)

type ExtrudeFuncFactory func(height float64, scale v2.Vec) sdf.ExtrudeFunc

func CustomScaleExtrude3D(sdf2 sdf.SDF2, height float64, scale v2.Vec, factory ExtrudeFuncFactory) sdf.SDF3 {
	s := ExtrudeSDF3{}
	s.sdf = sdf2
	s.height = height / 2
	s.extrude = factory(height, scale)

	// work out the bounding box
	bb := sdf2.BoundingBox()
	bb = bb.Extend(sdf.Box2{Min: bb.Min.Mul(scale), Max: bb.Max.Mul(scale)})
	s.bb = sdf.Box3{Min: v3.Vec{X: bb.Min.X, Y: bb.Min.Y, Z: -s.height}, Max: v3.Vec{X: bb.Max.X, Y: bb.Max.Y, Z: s.height}}
	return &s
}

func CosExtrude(height float64, scale v2.Vec) sdf.ExtrudeFunc {
	h := math.Pi / height
	sc := scale.SubScalar(1)
	return func(p v3.Vec) v2.Vec {
		z := p.Z * h
		if z < 0 {
			z = -z
		}
		r := math.Cos(z)
		s := sc.MulScalar(r).AddScalar(1)
		return v2.Vec{X: p.X, Y: p.Y}.Div(s)
	}
}

func CircExtrude(height float64, scale v2.Vec) sdf.ExtrudeFunc {
	h := 2 / height
	return func(p v3.Vec) v2.Vec {
		z := normZ(p.Z * h)
		r := math.Sqrt(1 - z*z)
		s := scale.SubScalar(1).MulScalar(r).AddScalar(1)
		return v2.Vec{X: p.X, Y: p.Y}.Div(s)
	}
}

func QuadExtrude(height float64, scale v2.Vec) sdf.ExtrudeFunc {
	h := 2 / height
	return func(p v3.Vec) v2.Vec {
		z := normZ(p.Z * h)
		r := -z*z + 1
		s := scale.SubScalar(1).MulScalar(r).AddScalar(1)
		return v2.Vec{X: p.X, Y: p.Y}.Div(s)
	}
}

func CubicExtrude(height float64, scale v2.Vec) sdf.ExtrudeFunc {
	h := 2 / height
	return func(p v3.Vec) v2.Vec {
		z := normZ(p.Z * h)
		r := -z*z*z + 1
		s := scale.SubScalar(1).MulScalar(r).AddScalar(1)
		return v2.Vec{X: p.X, Y: p.Y}.Div(s)
	}
}

func QuartExtrude(height float64, scale v2.Vec) sdf.ExtrudeFunc {
	h := 2 / height
	return func(p v3.Vec) v2.Vec {
		z := normZ(p.Z * h)
		r := -z*z*z*z + 1
		s := scale.SubScalar(1).MulScalar(r).AddScalar(1)
		return v2.Vec{X: p.X, Y: p.Y}.Div(s)
	}
}

func QuintExtrude(height float64, scale v2.Vec) sdf.ExtrudeFunc {
	h := 2 / height
	return func(p v3.Vec) v2.Vec {
		z := normZ(p.Z * h)
		r := -z*z*z*z*z + 1
		s := scale.SubScalar(1).MulScalar(r).AddScalar(1)
		return v2.Vec{X: p.X, Y: p.Y}.Div(s)
	}
}

func ExpoExtrude(height float64, scale v2.Vec) sdf.ExtrudeFunc {
	h := 2 / height
	return func(p v3.Vec) v2.Vec {
		z := normZ(p.Z * h)
		r := 1 - math.Exp2(10*(z-1))
		s := scale.SubScalar(1).MulScalar(r).AddScalar(1)
		return v2.Vec{X: p.X, Y: p.Y}.Div(s)
	}
}

func normZ(z float64) float64 {
	if z < 0 {
		z = -z
	}
	if z > 1 {
		z = 1
	}
	return z
}
