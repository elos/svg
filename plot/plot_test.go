package plot_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/elos/svg"
	"github.com/elos/svg/plot"
)

type p struct{ x, y float64 }

func (p p) X() float64 { return p.x }
func (p p) Y() float64 { return p.y }

func TestReverse(t *testing.T) {
	cases := []struct {
		have []plot.Point
		want []plot.Point
	}{
		{
			have: []plot.Point{p{1, 1}, p{2, 2}, p{3, 3}},
			want: []plot.Point{p{3, 3}, p{2, 2}, p{1, 1}},
		},
		{
			have: []plot.Point{p{1, 1}, p{2, 2}, p{3, 3}, p{4, 4}},
			want: []plot.Point{p{4, 4}, p{3, 3}, p{2, 2}, p{1, 1}},
		},
	}

	for _, c := range cases {
		if got := plot.Reverse(c.have); !reflect.DeepEqual(got, c.want) {
			t.Fatalf("plot.Reverse(%v): got\n\t*%vwant\n\t*%v", got, c.want)
		}
	}
}

func TestExtrema(t *testing.T) {
	cases := []struct {
		have []plot.Point
		want struct {
			minX, minY, maxX, maxY float64
		}
	}{
		{
			have: []plot.Point{p{1, 1}},
			want: struct{ minX, minY, maxX, maxY float64 }{
				minX: 1, minY: 1, maxX: 1, maxY: 1,
			},
		},
		{
			have: []plot.Point{p{1, 1}, p{-1, 10}, p{0, 20}, p{-10, 30}, p{100, 5}},
			want: struct{ minX, minY, maxX, maxY float64 }{
				minX: -10, minY: 1, maxX: 100, maxY: 30,
			},
		},
	}

	for _, c := range cases {
		t.Logf("Have: %v", c.have)
		minX, minY, maxX, maxY := plot.Extrema(c.have)

		if got, want := minX, c.want.minX; got != want {
			t.Errorf("\tminX: got %f, want %f", got, want)
		}
		if got, want := minY, c.want.minY; got != want {
			t.Errorf("\tminY: got %f, want %f", got, want)
		}
		if got, want := maxX, c.want.maxX; got != want {
			t.Errorf("\tmaxX: got %f, want %f", got, want)
		}
		if got, want := maxY, c.want.maxY; got != want {
			t.Errorf("\tmaxY: got %f, want %f", got, want)
		}
	}
}

func TestSample(t *testing.T) {
	cases := []struct {
		have []plot.Point
		n    int
		want []plot.Point
	}{
		{
			have: []plot.Point{p{1, 1}},
			n:    1,
			want: []plot.Point{p{1, 1}},
		},
		{
			have: []plot.Point{p{1, 1}, p{2, 2}, p{3, 3}},
			n:    2,
			want: []plot.Point{p{1, 1}, p{3, 3}},
		},
		{
			have: []plot.Point{p{1, 1}, p{2, 2}, p{3, 3}},
			n:    3,
			want: []plot.Point{p{1, 1}},
		},
		{
			have: []plot.Point{p{1, 1}, p{2, 2}, p{3, 3}},
			n:    5,
			want: []plot.Point{p{1, 1}},
		},
	}

	for _, c := range cases {
		t.Logf("c.have: %v", c.have)
		if got, want := plot.Sample(c.have, c.n), c.want; !reflect.DeepEqual(got, want) {
			t.Errorf("\tplot.Sample(c.have, %d): got:\n\t%bwant:\n\t%v", c.n, got, want)
		}
	}
}

func TestLine(t *testing.T) {
	cases := []struct {
		points []plot.Point
		width  float64
		height float64
		want   *svg.SVG
	}{
		{
			points: []plot.Point{p{1, 1}, p{2, 2}, p{3, 3}},
			width:  200,
			height: 100,
			want: &svg.SVG{
				Width:  200,
				Height: 100,
				Children: []svg.Encoder{
					&svg.Path{
						D: []*svg.PathCommand{
							{
								Directive: svg.MoveTo,
								Point:     &svg.Point{X: 0, Y: 0},
							},
							{
								Directive: svg.LineTo,
								Point:     &svg.Point{X: 100, Y: 50},
							},
							{
								Directive: svg.LineTo,
								Point:     &svg.Point{X: 200, Y: 100},
							},
						},
						Presentation: &svg.Presentation{
							Stroke:      "red",
							StrokeWidth: "0.8",
						},
					},
					// horizontal
					&svg.Line{X1: 0, Y1: 20, X2: 200, Y2: 20, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 0, Y1: 40, X2: 200, Y2: 40, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 0, Y1: 60, X2: 200, Y2: 60, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 0, Y1: 80, X2: 200, Y2: 80, Presentation: plot.AxesPresentation},
					// vertical
					&svg.Line{X1: 20, Y1: 0, X2: 20, Y2: 100, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 40, Y1: 0, X2: 40, Y2: 100, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 60, Y1: 0, X2: 60, Y2: 100, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 80, Y1: 0, X2: 80, Y2: 100, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 100, Y1: 0, X2: 100, Y2: 100, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 120, Y1: 0, X2: 120, Y2: 100, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 140, Y1: 0, X2: 140, Y2: 100, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 160, Y1: 0, X2: 160, Y2: 100, Presentation: plot.AxesPresentation},
					&svg.Line{X1: 180, Y1: 0, X2: 180, Y2: 100, Presentation: plot.AxesPresentation},
				},
				Presentation: &svg.Presentation{
					Stroke:      "black",
					Fill:        "white",
					StrokeWidth: "0.7",
				},
			},
		},
	}

	for _, c := range cases {
		t.Logf("c.points: %v", c.points)
		if got, want := plot.Line(c.points, c.width, c.height), c.want; !reflect.DeepEqual(got, want) {
			b := new(bytes.Buffer)
			got.Encode(b)
			t.Logf("Got:\n%s", b.String())
			b.Reset()
			want.Encode(b)
			t.Logf("Want:\n%s", b.String())
			t.Errorf("plot.Line(c.points, %f, %f):\n\tgot\n\t\t%+v\n\twant\n\t\t%+v", c.width, c.height, got, want)
		}
	}

}

func TestAxes(t *testing.T) {
	cases := []struct {
		xMin, yMin, xMax, yMax, xUnit, yUnit, height, width float64
		want                                                []svg.Encoder
	}{
		{
			xMin: -10, yMin: -10, xMax: 10, yMax: 10, xUnit: 1, yUnit: 1, width: 40, height: 40,
			want: []svg.Encoder{
				&svg.Line{
					X1: 0, Y1: 20, X2: 40, Y2: 20, Presentation: plot.AxesPresentation,
				},
				&svg.Line{
					X1: 20, Y1: 0, X2: 20, Y2: 40, Presentation: plot.AxesPresentation,
				},
			},
		},
	}

	buf := new(bytes.Buffer)
	for _, c := range cases {
		if got, want := plot.Axes(c.xMin, c.yMin, c.xMax, c.yMax, c.xUnit, c.yUnit, c.width, c.height), c.want; !reflect.DeepEqual(got, want) {
			buf.Reset()
			for _, e := range got {
				e.Encode(buf)
			}
			t.Logf("got:\n%s", buf.String())
			buf.Reset()
			for _, e := range want {
				e.Encode(buf)
			}
			t.Logf("want:\n%s", buf.String())
			t.Errorf("plot.Axes(%f, %f, %f, %f, %f, %f): got != want", c.xMin, c.yMin, c.xMax, c.yMax, c.height, c.width)
		}
	}
}
