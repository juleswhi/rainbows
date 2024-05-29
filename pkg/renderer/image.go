package renderer

type Image struct {
	width  int
	height int
}

func NewImage(w, h int) *Image {
	return &Image{
		width:  w,
		height: h,
	}
}
