package svg

import (
	"fmt"
	"io"
)

type Encoder interface {
	Encode(io.Writer)
}

type Presentation struct {
	Color,
	Fill,
	Opacity,
	Stroke,
	StrokeDashArray,
	StrokeWidth string
}

func (p *Presentation) Encode(w io.Writer) {
	if p.Color != "" {
		fmt.Fprintf(w, " color=%q", p.Color)
	}
	if p.Fill != "" {
		fmt.Fprintf(w, " fill=%q", p.Fill)
	}
	if p.Opacity != "" {
		fmt.Fprintf(w, " opacity=%q", p.Opacity)
	}
	if p.Stroke != "" {
		fmt.Fprintf(w, " stroke=%q", p.Stroke)
	}
	if p.StrokeDashArray != "" {
		fmt.Fprintf(w, " stroke-dasharray=%q", p.StrokeDashArray)
	}
	if p.StrokeWidth != "" {
		fmt.Fprintf(w, " stroke-width=%q", p.StrokeWidth)
	}
}

// --- Anchor {{{

// Anchor defines the structure of an anchor element.
type Anchor struct {
	Show, Actuate, HRef, Target string
	Children                    []Encoder
}

// Encode writes the encoding of an anchor element.
func (a *Anchor) Encode(w io.Writer) {
	fmt.Fprintf(
		w,
		`<a xlink:show="%s" xlink:actuate="%s" xlink:href="%s" xlink:target="%s">`,
		a.Show, a.Actuate, a.HRef, a.Target,
	)
	for _, c := range a.Children {
		c.Encode(w)
	}
	fmt.Fprint(w, "</a>")
}

// --- }}}

// --- Circle {{{

// Circle defines the structure of a circle element.
type Circle struct {
	*Presentation
	CX, CY, R float64
}

// Encode writes the encoding of a circle element.
func (c *Circle) Encode(w io.Writer) {
	fmt.Fprintf(
		w,
		`<circle cx="%.2f" cy="%.2f" r="%.2f" />`,
		c.CX, c.CY, c.R,
	)
}

// --- }}}

// --- Ellipse {{{

// Ellipse defines the structure of an ellipse element.
type Ellipse struct {
	*Presentation
	CX, CY, RX, RY float64
}

// Encode writes the encoding of an ellipse element.
func (e *Ellipse) Encode(w io.Writer) {
	fmt.Fprintf(
		w,
		`<ellipse cx="%.2f" cy="%.2f" rx="%.2f" ry="%.2f" />`,
		e.CX, e.CY, e.RX, e.RY,
	)
}

// --- }}}

// --- Group {{{

// Group defines the structure of a group element.
type Group struct {
	*Presentation
	Children []Encoder
}

// Encode writes the encoding of a group element.
func (g *Group) Encode(w io.Writer) {
	fmt.Fprint(w, "<g>")
	for _, c := range g.Children {
		c.Encode(w)
	}
	fmt.Fprintf(w, "</g>")
}

// --- }}}

// --- Image {{{

// Image defines the structure of an image element.
type Image struct {
	X, Y, Width, Height float64
	HRef                string
}

// Encode writes the encoding of an image element.
func (i *Image) Encode(w io.Writer) {
	fmt.Fprintf(
		w,
		`<image xlink:href="%s" x="%.2f" y="%.2f" width="%.2f" height="%.2f" />`,
		i.HRef, i.X, i.Y, i.Width, i.Height,
	)
}

// --- }}}

// --- Line {{{

// Line defines the structure of a line element.
type Line struct {
	X1, Y1, X2, Y2 float64
	*Presentation
}

// Encode writes the encoding of a line element.
func (l *Line) Encode(w io.Writer) {
	fmt.Fprintf(
		w,
		`<line x1="%.2f" y1="%.2f" x2="%.2f" y2="%.2f"`,
		l.X1, l.Y1, l.X2, l.Y2,
	)
	if l.Presentation != nil {
		l.Presentation.Encode(w)
	}
	fmt.Fprint(w, `/>`)
}

// --- }}}

// --- Path {{{

// Path defines the structure of a path element.
type Path struct {
	D []*PathCommand
	*Presentation
}

// Encode writes the encoding of a path.
func (p *Path) Encode(w io.Writer) {
	fmt.Fprint(w, `<path d="`)
	for _, c := range p.D {
		c.Encode(w)
	}
	fmt.Fprint(w, `"`)
	if p.Presentation != nil {
		p.Presentation.Encode(w)
	}
	fmt.Fprint(w, ` />`)
}

// PathCommand defines the structure of a command, in
// the path "d" attribute of a Path.
//
// Note: All directives include a point except the ClosePath
// directive.
type PathCommand struct {
	Directive PathDirective
	*Point
}

// Encode writes the encoding of a PathCommand.
func (pc *PathCommand) Encode(w io.Writer) {
	switch pc.Directive {
	case ClosePath:
		fmt.Fprint(w, "Z")
	default:
		fmt.Fprintf(w, "%s%.2f %.2f ", pc.Directive, pc.Point.X, pc.Point.Y)
	}
}

// A PathDirective is the single letter directive defining the type
// of a PathCommand.
type PathDirective string

const (
	MoveTo                     PathDirective = "M"
	LineTo                                   = "L"
	HorizontalLineTo                         = "H"
	VerticalLineTo                           = "V"
	CurveTo                                  = "C"
	SmoothCurveTo                            = "C"
	QuadraticBezierCurve                     = "Q"
	SmoothQuadraticBezierCurve               = "T"
	EllipticalArc                            = "A"
	ClosePath                                = "Z"
)

// A Point defines the structure of a 2D coordinate.
type Point struct {
	X, Y float64
}

// --- }}}

type Polygon struct {
	Points []*Point
}

func (p *Polygon) Encode(w io.Writer) {
	fmt.Fprintf(w, `<polygon points="`)
	for _, point := range p.Points {
		fmt.Fprintf(w, "%f,%f", point.X, point.Y)
	}
	fmt.Fprintf(w, `" />`)
}

type Polyline struct {
	Points []*Point
}

func (p *Polyline) Encode(w io.Writer) {
	fmt.Fprint(w, `<polyline points="`)
	for _, point := range p.Points {
		fmt.Fprintf(w, "%f,%f", point.X, point.Y)
	}
	fmt.Fprintf(w, `" />`)
}

type Rect struct {
	X, Y, RX, RY, Width, Height float64
}

func (r *Rect) Encode(w io.Writer) {
	fmt.Fprintf(
		w,
		`<rect x="%f" y="%f" rx="%f" ry="%f" width="%f" height="%f" />`,
		r.X, r.Y, r.RX, r.RY, r.Width, r.Height,
	)
}

type SVG struct {
	Width, Height float64
	Children      []Encoder
	Presentation  *Presentation
}

func (svg *SVG) Encode(w io.Writer) {
	fmt.Fprintf(w, `<svg xmlns="http://www.w3.org/2000/svg" width="%.2f" height="%.2f"`, svg.Width, svg.Height)
	svg.Presentation.Encode(w)
	fmt.Fprint(w, ` >`)
	for _, c := range svg.Children {
		c.Encode(w)
	}
	fmt.Fprintf(w, "</svg>")
}

type Text struct {
	Content string
	*Point
}

func (t *Text) Encode(w io.Writer) {
	fmt.Fprintf(w, `<text x="%.2f" y="%.2f">%s</text>`, t.Point.X, t.Point.Y, t.Content)
}
