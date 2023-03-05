package svg

import "fmt"

const (
	mpi = 25.4
)

var (
	Default = DPI{
		ppi: mpi,
		mpd: 1,
	}
)

type DPI struct {
	ppi float64
	mpd float64
}

func NewDPI(dpi float64) DPI {
	return DPI{
		ppi: dpi,
		mpd: mpi / dpi,
	}
}

func NewDPIWithSize(pixel, mm float64) DPI {
	mpd := mm / pixel
	return DPI{
		ppi: mpi / mpd,
		mpd: mpd,
	}
}

func (d DPI) DPI() float64 {
	return d.ppi
}

// MM transforms pixel to mm.
func (d DPI) MM(pixel float64) float64 {
	return pixel * d.mpd
}

// Pixel transforms mm to pixel.
func (d DPI) Pixel(mm float64) float64 {
	return mm / d.mpd
}

func (d DPI) String() string {
	return fmt.Sprintf("%f", d.ppi)
}
