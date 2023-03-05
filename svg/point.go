package svg

var (
	ZeroPoint = NewPoint(0, 0)
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func (pt Point) Add(point Point) Point {
	return NewPoint(
		pt.X+point.X,
		pt.Y+point.Y,
	)
}

// Matrix transforms coordinates from a previous coordinate system into a new coordinate system by following matrix equalities:
// | A C E |   | X |   | AX+CY+E |
// | B D F | * | Y | = | BX+DY+F |
// | 0 0 1 |   | 1 |   | 1       |
func (pt Point) Matrix(m *Matrix) Point {
	if m == nil {
		return pt
	}
	return NewPoint(m.A*pt.X+m.C*pt.Y+m.E, m.B*pt.X+m.D*pt.Y+m.F)
}

func (pt Point) MM(dpi DPI) Point {
	return NewPoint(dpi.MM(pt.X), dpi.MM(pt.Y))
}
