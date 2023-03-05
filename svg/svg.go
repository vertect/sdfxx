package svg

import (
	"encoding/xml"
	"fmt"
	"io"
)

type SVG struct {
	Width    string     `xml:"width,attr"`
	Height   string     `xml:"height,attr"`
	ViewBox  string     `xml:"viewBox,attr"`
	Fill     string     `xml:"fill,attr"`
	XMLNS    string     `xml:"xmlns,attr"`
	Graphics []*Graphic `xml:"g"`

	dpi DPI
}

type Option struct {
	DPI DPI
}

func New(r io.Reader, option *Option) (*SVG, error) {
	var svg SVG
	if err := xml.NewDecoder(r).Decode(&svg); err != nil {
		return nil, err
	}
	if option != nil {
		svg.setDPI(option.DPI)
	} else {
		svg.setDPI(Default)
	}
	return &svg, nil
}

func (svg *SVG) Find(id string) (*Graphic, error) {
	for _, g := range svg.Graphics {
		f := g.Find(id)
		if f != nil {
			return f, nil
		}
	}
	return nil, NotFoundError(fmt.Sprintf("%s not found", id))
}

func (svg *SVG) setDPI(dpi DPI) {
	svg.dpi = dpi
	for _, g := range svg.Graphics {
		g.setDPI(dpi)
	}
}

type NotFoundError string

func (err NotFoundError) Error() string {
	return string(err)
}
