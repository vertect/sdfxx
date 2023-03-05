package svg

import (
	"encoding/xml"
	"regexp"
	"strconv"

	"github.com/deadsy/sdfx/sdf"
)

type Graphic struct {
	ID     string
	Matrix *Matrix

	Paths    []*Path
	Rects    []*Rect
	Circles  []*Circle
	Graphics []*Graphic

	dpi DPI
}

type graphic struct {
	ID        string     `xml:"id,attr"`
	Transform string     `xml:"transform,attr"`
	Paths     []*Path    `xml:"path"`
	Rects     []*Rect    `xml:"rect"`
	Circles   []*Circle  `xml:"circle"`
	Graphics  []*Graphic `xml:"g"`
}

func (gr *Graphic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var g graphic
	if err := d.DecodeElement(&g, &start); err != nil {
		return err
	}
	matrix, err := parseTransform(g.Transform)
	if err != nil {
		return err
	}
	for _, g := range g.Graphics {
		g.applyMatrix(matrix)
	}
	*gr = Graphic{
		ID:       g.ID,
		Matrix:   matrix,
		Paths:    g.Paths,
		Rects:    g.Rects,
		Circles:  g.Circles,
		Graphics: g.Graphics,
	}
	return nil
}

var rMatrix = regexp.MustCompile(`matrix\(([-\d\.]+?)\s*,\s*([-\d\.]+?)\s*,\s*([-\d\.]+?)\s*,\s*([-\d\.]+?)\s*,\s*([-\d\.]+?)\s*,\s*([-\d\.]+?)\)`)

func parseTransform(transform string) (*Matrix, error) {
	result := rMatrix.FindStringSubmatch(transform)
	if len(result) == 0 {
		return nil, nil
	}
	elements := make([]float64, 6)
	for i := 0; i < 6; i++ {
		f, err := strconv.ParseFloat(result[i+1], 64)
		if err != nil {
			return nil, err
		}
		elements[i] = f
	}
	return NewMatrix(elements[0], elements[1], elements[2], elements[3], elements[4], elements[5]), nil
}

func (gr *Graphic) applyMatrix(matrix *Matrix) {
	gr.Matrix = matrix.Dot(gr.Matrix)
}

func (gr *Graphic) Find(id string) *Graphic {
	if gr.ID == id {
		return gr
	}
	for _, g := range gr.Graphics {
		f := g.Find(id)
		if f != nil {
			return f
		}
	}
	return nil
}

func (gr *Graphic) Draw() (sdf.SDF2, error) {
	var polygons []sdf.SDF2
	for _, path := range gr.Paths {
		bs, err := path.ToBeziers(gr.Matrix)
		if err != nil {
			return nil, err
		}
		for _, b := range bs {
			p, err := b.Polygon()
			if err != nil {
				return nil, err
			}
			polygon, err := sdf.Polygon2D(p.Vertices())
			if err != nil {
				return nil, err
			}
			polygons = append(polygons, polygon)
		}
	}
	for _, r := range gr.Rects {
		p, err := r.Draw()
		if err != nil {
			return nil, err
		}
		polygons = append(polygons, p)
	}
	for _, c := range gr.Circles {
		p, err := c.Draw()
		if err != nil {
			return nil, err
		}
		polygons = append(polygons, p)
	}
	for _, g := range gr.Graphics {
		p, err := g.Draw()
		if err != nil {
			return nil, err
		}
		polygons = append(polygons, p)
	}
	return sdf.Union2D(polygons...), nil
}

func (gr *Graphic) setDPI(dpi DPI) {
	gr.dpi = dpi
	for _, p := range gr.Paths {
		p.setDPI(dpi)
	}
	for _, r := range gr.Rects {
		r.setDPI(dpi)
	}
	for _, c := range gr.Circles {
		c.setDPI(dpi)
	}
	for _, g := range gr.Graphics {
		g.setDPI(dpi)
	}
}
