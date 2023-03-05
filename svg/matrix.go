package svg

// Matrix is equivalent to applying the transform matrix:
// | A C E |
// | B D F |
// | 0 0 1 |
type Matrix struct {
	A float64
	B float64
	C float64
	D float64
	E float64
	F float64
}

func NewMatrix(a, b, c, d, e, f float64) *Matrix {
	return &Matrix{a, b, c, d, e, f}
}

// Dot multiply 2 matrix:
// | A C E |   | a c e |   | Aa+Cb Ac+Cd Ae+Cf+E |
// | B D F | * | b d f | = | Ba+Db Bc+Dd Be+Df+F |
// | 0 0 1 |   | 0 0 1 |   | 0     0     1       |
func (m *Matrix) Dot(matrix *Matrix) *Matrix {
	if matrix == nil {
		return m
	}
	if m == nil {
		return matrix
	}
	A := m.A
	B := m.B
	C := m.C
	D := m.D
	E := m.E
	F := m.F
	a := matrix.A
	b := matrix.B
	c := matrix.C
	d := matrix.D
	e := matrix.E
	f := matrix.F
	return NewMatrix(
		A*a+C*b, B*a+D*b, A*c+C*d, B*c+D*d, A*e+C*f+E, B*e+D*f+F,
	)
}
