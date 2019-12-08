package plot

import (
	"math"

	"github.com/elos/svg"
)

type Point interface {
	X() float64
	Y() float64
}

// Reverse reverses the points.
func Reverse(points []Point) []Point {
	for i, j := 0, len(points)-1; i < j; i, j = i+1, j-1 {
		points[i], points[j] = points[j], points[i]
	}
	return points
}

// Extrema finds the extrema of the points.
//
// Use extrema to find the coordinates which will define the size
// of the chart.
func Extrema(points []Point) (minX, minY, maxX, maxY float64) {
	minX, minY = points[0].X(), points[0].Y()
	maxX, maxY = points[0].X(), points[0].Y()

	for i := 1; i < len(points); i++ {
		x, y := points[i].X(), points[i].Y()

		if x < minX {
			minX = x
		}

		if y < minY {
			minY = y
		}

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}
	}

	return
}

// Sample samples every n points.
//
// Use sample to reduce the number of data points
// for dense (sub-second) frequency data.
func Sample(points []Point, n int) []Point {
	s := make([]Point, int(math.Ceil(float64(len(points))/float64(n))))
	for i := 0; i < len(points); i += n {
		s[i/n] = points[i]
	}
	return s
}

// Line produces a line plot of the points.
func Line(points []Point, width, height float64) *svg.SVG {
	xMin, yMin, xMax, yMax := Extrema(points)
	xUnit, yUnit := width/(xMax-xMin), height/(yMax-yMin)
	s := &svg.SVG{
		Width:  width,
		Height: height,
		Presentation: &svg.Presentation{
			Stroke:      "black",
			Fill:        "white",
			StrokeWidth: "0.7",
		},
	}
	p := &svg.Path{
		Presentation: &svg.Presentation{
			Stroke:      "red",
			StrokeWidth: "0.8",
		},
	}

	p.D = append(p.D, &svg.PathCommand{
		Directive: svg.MoveTo,
		Point: &svg.Point{
			X: (points[0].X() - xMin) * xUnit,
			Y: (points[0].Y() - yMin) * yUnit,
		},
	})

	for i := 1; i < len(points); i++ {
		p.D = append(p.D, &svg.PathCommand{
			Directive: svg.LineTo,
			Point: &svg.Point{
				X: (points[i].X() - xMin) * xUnit,
				Y: (points[i].Y() - yMin) * yUnit,
			},
		})
	}

	s.Children = append(s.Children, p)
	s.Children = append(s.Children, Axes(xMin, yMin, xMax, yMax, xUnit, yUnit, width, height)...)

	return s
}

// pixels
const axisStride = 20.0

var AxesPresentation = &svg.Presentation{
	Opacity: "0.1",
}

func Axes(xMin, yMin, xMax, yMax, xUnit, yUnit, width, height float64) []svg.Encoder {
	// horizontal
	es := make([]svg.Encoder, 0)
	for y := axisStride; y < height; y += axisStride {
		es = append(es, &svg.Line{
			X1:           0,
			Y1:           y,
			Y2:           y,
			X2:           width,
			Presentation: AxesPresentation,
		})
	}

	// vertical
	for x := axisStride; x < width; x += axisStride {
		es = append(es, &svg.Line{
			X1:           x,
			Y1:           0,
			X2:           x,
			Y2:           height,
			Presentation: AxesPresentation,
		})
	}

	return es
}
