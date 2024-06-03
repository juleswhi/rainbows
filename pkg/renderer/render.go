package renderer

import (
	"math"
)

type Camera *V3

func newCamera() *V3 {
	return newV3(0, 0, 2)
}

type Plane struct {
	X1 *V3
	X2 *V3
	X3 *V3
	X4 *V3
}

type Colour struct {
	R float32
	G float32
	B float32
}

func newColour(r, g, b float32) *Colour {
	return &Colour{
		R: r,
		G: g,
		B: b,
	}
}

type Sphere struct {
	C      *V3
	R      float32
	Colour *Colour
}

func newPlane() Plane {
	return Plane{
		X1: newV3(-1.28, 0.86, -0.5),
		X2: newV3(1.28, 0.86, -0.5),
		X3: newV3(-1.28, -0.86, -0.5),
		X4: newV3(1.28, -0.86, -0.5),
	}
}

func newSphere(center *V3, radius float32, colour *Colour) Sphere {
	return Sphere{
		C:      center,
		R:      radius,
		Colour: colour,
	}
}

func newSpheres() []Sphere {
	return []Sphere{
		newSphere(newV3(0.2, -0.1, -1), 0.5, newColour(0, 1, 0)),
		newSphere(newV3(1.2, -0.5, -1.75), 0.4, newColour(1, 1, 1)),
	}
}

func (s *Sphere) Intersection(ray Ray) float64 {
	cp := ray.Origin.Take(s.C)

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * cp.Dot(ray.Direction)
	c := cp.Dot(cp) - s.R*s.R

	disc := b*b - 4*a*c

	if disc < 0 {
		return 0
	}

	sqrt := math.Sqrt(float64(disc))

	var ts []float64

	sub := (float64(-b) - sqrt) / (2 * float64(a))

	if sub >= 0 {
		ts = append(ts, sub)
	}

	add := (float64(-b) + sqrt) / (2 * float64(a))

	if add >= 0 {
		ts = append(ts, add)
	}

	if len(ts) == 0 {
		return 0
	}

	return minfloat(ts)
}

func minfloat(floats []float64) (m float64) {
	if len(floats) > 0 {
		m = floats[0]
	}

	for i := 1; i < len(floats); i++ {
		if floats[i] < m {
			m = floats[i]
		}
	}

	return
}

func RayCast() []Colour {
	image := NewImage(256, 192)
	plane := newPlane()
	rays := []Ray{}
	camera := newCamera()

	for y := range image.height {
		for x := range image.width {
			a := (float32(x) / float32(image.width))
			B := (float32(y) / float32(image.height))

			t := plane.X1.Scale(1 - a).Add(plane.X2.Scale(a))
			b := plane.X3.Scale(1 - a).Add(plane.X4.Scale(a))

			p := t.Scale(1 - B).Add(b.Scale(B))
			rays = append(rays, *newRay(p, p.Take(camera)))
		}
	}

	var maxX float32 = 0
	var minX float32 = 0
	var maxY float32 = 0
	var minY float32 = 0

	for _, ray := range rays {
		if ray.Direction.Y > maxY {
			maxY = ray.Direction.Y
		}
		if ray.Direction.Y < minY {
			minY = ray.Direction.Y
		}

		if ray.Direction.X > maxX {
			maxX = ray.Direction.X
		}
		if ray.Direction.X < minX {
			minX = ray.Direction.X
		}
	}

	cols := []Colour{}
	spheres := newSpheres()

	for _, ray := range rays {
		var intersections []float64

		for _, sphere := range spheres {
			inter := sphere.Intersection(ray)
			intersections = append(intersections, inter)
		}

		if allZero(intersections) {
			col := newColour(
				0, 0, 0,
			)
			cols = append(cols, *col)
			continue
		}

		minIdx := idxNonZero(intersections)

		sp := spheres[minIdx]

		col := newColour(
			sp.Colour.R*255,
			sp.Colour.G*255,
			sp.Colour.B*255,
		)

		cols = append(cols, *col)
	}

	return cols
}

func allZero(nums []float64) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func idxNonZero(nums []float64) int {
	minVal := math.MaxFloat64
	minIdx := 0

	for i, num := range nums {
		if num == 0 {
			continue
		}
		if num < minVal {
			minVal = num
			minIdx = i
		}
	}
	return minIdx
}

func normalize(value float32, minX, maxX, newMin, newMax float32) float32 {
	oldRange := maxX - minX
	newRange := newMax - newMin
	return (((value - minX) * newRange) / oldRange) + newMin
}
