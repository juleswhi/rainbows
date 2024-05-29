package renderer

import (
	"github.com/charmbracelet/log"
)

type Camera *V3

func newCamera() *V3 {
	return newV3(0, 0, -1)
}

type Plane struct {
	X1 *V3
	X2 *V3
	X3 *V3
	X4 *V3
}

func newPlane() Plane {
	return Plane{
		X1: newV3(1, 0.75, 0),
		X2: newV3(-1, 0.75, 0),
		X3: newV3(1, -0.75, 0),
		X4: newV3(-1, -0.75, 0),
	}
}

func RayCast() []Ray {
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

	newRays := []Ray{}

	for _, ray := range rays {
		newray := ray
		newray.Direction.X = normalize(ray.Direction.X, minX, maxX, 0, 255)
		newray.Direction.Y = normalize(ray.Direction.Y, minX, maxX, 0, 255)
		newray.Direction.Z = normalize(ray.Direction.Z, minX, maxX, 0, 255)
		newRays = append(newRays, newray)
	}

	log.Info("X", "Max", maxX, "Min", minX)
	log.Info("Y", "Max", maxY, "Min", minY)

	// log.Info("X * 255", "Max", maxX*255, "Min", minX*255)
	// log.Info("Y * 255", "Max", maxY*255, "Min", minY*255)

	return newRays
}

func normalize(value float32, minX, maxX, newMin, newMax float32) float32 {
	oldRange := maxX - minX
	newRange := newMax - newMin
	return (((value - minX) * newRange) / oldRange) + newMin
}
