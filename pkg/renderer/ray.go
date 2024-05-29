package renderer

type Ray struct {
    Origin *V3
    Direction *V3
}

func newRay(p, d *V3) *Ray {
    return &Ray {
        Origin: p,
        Direction: d,
    }
}
