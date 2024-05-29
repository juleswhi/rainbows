package renderer

type V3 struct {
	X float32
	Y float32
	Z float32
}

func newV3(x, y, z float32) *V3 {
	return &V3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (vec1 *V3) Add(vec2 *V3) *V3 {
    return newV3(vec1.X + vec2.X, vec1.Y + vec2.Y, vec1.Z + vec2.Z)
}

func (vec1 *V3) Take(vec2 *V3) *V3 {
    return newV3(vec1.X - vec2.X, vec1.Y - vec2.Y, vec1.Z - vec2.Z)
}


func (vec1 *V3) Scale(scale float32) *V3 {
    return &V3 {
        X: vec1.X * scale,
        Y: vec1.Y * scale,
        Z: vec1.Z * scale,
    }
}
