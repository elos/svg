package svg_test

import (
	"bytes"
	"testing"

	"github.com/elos/svg"
)

func TestAnchor(t *testing.T) {
	b := new(bytes.Buffer)

	cases := []struct {
		anchor *svg.Anchor
		want   []byte
	}{
		{
			anchor: &svg.Anchor{
				Show:    "show",
				Actuate: "actuate",
				HRef:    "href",
				Target:  "target",
			},
			want: []byte(`<a xlink:show="show" xlink:actuate="actuate" xlink:href="href" xlink:target="target"></a>`),
		},
	}

	for _, c := range cases {
		c.anchor.Encode(b)

		if got := b.Bytes(); !bytes.Equal(got, c.want) {
			t.Errorf("anchor.Encode: got:\n%swant:\n%s", got, c.want)
		}

		b.Reset()
	}
}

func TestCircle(t *testing.T) {
	b := new(bytes.Buffer)

	cases := []struct {
		circle *svg.Circle
		want   []byte
	}{
		{
			circle: &svg.Circle{
				CX: 4.5,
				CY: 5.4,
				R:  6,
			},
			want: []byte(`<circle cx="4.50" cy="5.40" r="6.00" />`),
		},
	}

	for _, c := range cases {
		c.circle.Encode(b)

		if got := b.Bytes(); !bytes.Equal(got, c.want) {
			t.Errorf("circle.Encode: got:\n%swant:\n%s", got, c.want)
		}

		b.Reset()
	}
}

func TestEllipse(t *testing.T) {
	b := new(bytes.Buffer)

	cases := []struct {
		ellipse *svg.Ellipse
		want    []byte
	}{
		{
			ellipse: &svg.Ellipse{
				CX: 4,
				CY: 5,
				RX: 4.5,
				RY: 5.5,
			},
			want: []byte(`<ellipse cx="4.00" cy="5.00" rx="4.50" ry="5.50" />`),
		},
	}

	for _, c := range cases {
		c.ellipse.Encode(b)

		if got := b.Bytes(); !bytes.Equal(got, c.want) {
			t.Errorf("ellipse.Encode: got:\n%swant:\n%s", got, c.want)
		}

		b.Reset()
	}
}

func TestGroup(t *testing.T) {
	b := new(bytes.Buffer)

	cases := []struct {
		group *svg.Group
		want  []byte
	}{
		{
			group: &svg.Group{
				Children: []svg.Encoder{
					&svg.Circle{
						CX: 4.5,
						CY: 5.4,
						R:  6,
					},
				},
			},
			want: []byte(`<g><circle cx="4.50" cy="5.40" r="6.00" /></g>`),
		},
	}

	for _, c := range cases {
		c.group.Encode(b)

		if got := b.Bytes(); !bytes.Equal(got, c.want) {
			t.Errorf("group.Encode: got:\n%swant:\n%s", got, c.want)
		}

		b.Reset()
	}
}

func TestImage(t *testing.T) {
	b := new(bytes.Buffer)

	cases := []struct {
		image *svg.Image
		want  []byte
	}{
		{
			image: &svg.Image{
				X:      4,
				Y:      4,
				Height: 5,
				Width:  5,
				HRef:   "href",
			},
			want: []byte(`<image xlink:href="href" x="4.00" y="4.00" width="5.00" height="5.00" />`),
		},
	}

	for _, c := range cases {
		c.image.Encode(b)

		if got := b.Bytes(); !bytes.Equal(got, c.want) {
			t.Errorf("image.Encode: got:\n%swant:\n%s", got, c.want)
		}

		b.Reset()
	}
}

func TestLine(t *testing.T) {
	b := new(bytes.Buffer)

	cases := []struct {
		line *svg.Line
		want []byte
	}{
		{
			line: &svg.Line{
				X1: 4,
				Y1: 4,
				X2: 5,
				Y2: 5,
			},
			want: []byte(`<line x1="4.00" y1="4.00" x2="5.00" y2="5.00" />`),
		},
	}

	for _, c := range cases {
		c.line.Encode(b)

		if got := b.Bytes(); !bytes.Equal(got, c.want) {
			t.Errorf("line.Encode: got:\n%swant:\n%s", got, c.want)
		}

		b.Reset()
	}
}

func TestPath(t *testing.T) {
	b := new(bytes.Buffer)

	cases := []struct {
		path *svg.Path
		want []byte
	}{
		{
			path: &svg.Path{
				D: []*svg.PathCommand{
					{
						Directive: svg.MoveTo,
						Point: &svg.Point{
							X: 5, Y: 5,
						},
					},
					{
						Directive: svg.LineTo,
						Point: &svg.Point{
							X: 5, Y: 10,
						},
					},
					{
						Directive: svg.ClosePath,
					},
				},
			},
			want: []byte(`<path d="M5.00 5.00 L5.00 10.00 Z" />`),
		},
	}

	for _, c := range cases {
		c.path.Encode(b)

		if got := b.Bytes(); !bytes.Equal(got, c.want) {
			t.Errorf("path.Encode: got:\n%swant:\n%s", got, c.want)
		}

		b.Reset()
	}
}
