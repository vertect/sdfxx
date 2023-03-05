package svg

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/deadsy/sdfx/sdf"
	"github.com/pkg/errors"
)

type Path struct {
	Drawables      []Drawable
	Stroke         string
	StrokeLinecap  string
	StrokeLineJoin string

	dpi DPI
}

type path struct {
	Draw           string `xml:"d,attr"`
	Stroke         string `xml:"stroke,attr"`
	StrokeLinecap  string `xml:"stroke-linecap,attr"`
	StrokeLineJoin string `xml:"stroke-linejoin,attr"`
}

func (pa *Path) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var p path
	if err := d.DecodeElement(&p, &start); err != nil {
		return err
	}
	drawables, err := parseDraw(p.Draw)
	if err != nil {
		return err
	}
	*pa = Path{
		Drawables:      drawables,
		Stroke:         p.Stroke,
		StrokeLinecap:  p.StrokeLinecap,
		StrokeLineJoin: p.StrokeLineJoin,
	}
	return nil
}

func parseDraw(draw string) ([]Drawable, error) {
	var drawables []Drawable
	l := len(draw)
	for i := 0; i < l; i++ {
		switch draw[i] {
		case 'M':
			p, j, err := readPoint2D(draw, i+1, l)
			if err != nil {
				return nil, err
			}
			i = j
			drawables = append(drawables, NewMoveTo(p))
		case 'L':
			p, j, err := readPoint2D(draw, i+1, l)
			if err != nil {
				return nil, err
			}
			i = j
			drawables = append(drawables, NewLineTo(p))
		case 'C':
			start, j, err := readPoint2D(draw, i+1, l)
			if err != nil {
				return nil, err
			}
			mid, j, err := readPoint2D(draw, j+2, l)
			if err != nil {
				return nil, err
			}
			end, j, err := readPoint2D(draw, j+2, l)
			if err != nil {
				return nil, err
			}
			i = j
			drawables = append(drawables, NewCubicBezier(
				start,
				mid,
				end,
			))
		case 'Z':
			drawables = append(drawables, NewClosePath())
		}
	}
	return drawables, nil
}

func (pa *Path) ToBeziers(m *Matrix) ([]*sdf.Bezier, error) {
	b := sdf.NewBezier()
	for _, d := range pa.Drawables {
		d.Matrix(m)
		d.MM(pa.dpi)
		d.DrawTo(b)
	}
	return []*sdf.Bezier{b}, nil
}

func (p *Path) setDPI(dpi DPI) {
	p.dpi = dpi
}

func readPoint2D(d string, start int, last int) (Point, int, error) {
	x, j, err := readFloat64(d, start, last)
	if err != nil {
		return ZeroPoint, j, err
	}
	y, j, err := readFloat64(d, j+2, last)
	if err != nil {
		return ZeroPoint, j, err
	}
	return NewPoint(x, y), j, err
}

func readFloat64(d string, start int, last int) (float64, int, error) {
	n := ""
	i := start
	for ; i < last; i++ {
		c := d[i]
		switch c {
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			n = n + string(c)
		default:
			f, err := strconv.ParseFloat(n, 64)
			return f, i - 1, errors.Wrap(err, fmt.Sprintf("fail to parse float %s", n))
		}
	}
	f, err := strconv.ParseFloat(n, 64)
	return f, i, err
}

type Drawable interface {
	Matrix(matrix *Matrix)
	MM(dpi DPI)
	DrawTo(b *sdf.Bezier)
}

type MoveTo struct {
	Point Point
}

func NewMoveTo(p Point) *MoveTo {
	return &MoveTo{
		p,
	}
}

func (m *MoveTo) Matrix(matrix *Matrix) {
	m.Point = m.Point.Matrix(matrix)
}

func (m *MoveTo) MM(dpi DPI) {
	m.Point = m.Point.MM(dpi)
}

func (m *MoveTo) DrawTo(b *sdf.Bezier) {
	b.Add(
		m.Point.X,
		m.Point.Y,
	)
}

type LineTo struct {
	Point Point
}

func NewLineTo(p Point) *LineTo {
	return &LineTo{
		p,
	}
}

func (l *LineTo) Matrix(matrix *Matrix) {
	l.Point = l.Point.Matrix(matrix)
}

func (l *LineTo) MM(dpi DPI) {
	l.Point = l.Point.MM(dpi)
}

func (l *LineTo) DrawTo(b *sdf.Bezier) {
	b.Add(
		l.Point.X,
		l.Point.Y,
	)
}

type CubicBezier struct {
	Start Point
	Mid   Point
	End   Point
}

func NewCubicBezier(start, mid, end Point) *CubicBezier {
	return &CubicBezier{start, mid, end}
}

func (cb *CubicBezier) Matrix(matrix *Matrix) {
	cb.Start = cb.Start.Matrix(matrix)
	cb.Mid = cb.Mid.Matrix(matrix)
	cb.End = cb.End.Matrix(matrix)
}

func (cb *CubicBezier) MM(dpi DPI) {
	cb.Start = cb.Start.MM(dpi)
	cb.Mid = cb.Mid.MM(dpi)
	cb.End = cb.End.MM(dpi)
}

func (cb *CubicBezier) DrawTo(b *sdf.Bezier) {
	b.Add(
		cb.Start.X,
		cb.Start.Y,
	).Mid()
	b.Add(
		cb.Mid.X,
		cb.Mid.Y,
	).Mid()
	b.Add(
		cb.End.X,
		cb.End.Y,
	)
}

type ClosePath struct{}

func NewClosePath() *ClosePath {
	return &ClosePath{}
}

func (m *ClosePath) Matrix(matrix *Matrix) {
}

func (m *ClosePath) MM(dpi DPI) {
}

func (m *ClosePath) DrawTo(b *sdf.Bezier) {
	b.Close()
}
