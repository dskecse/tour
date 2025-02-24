package main

import (
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

type Image struct {
    Width, Height int
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
    return i.Bounds().At(x, y)
}

func main() {
    m := Image{160, 40}
    pic.ShowImage(m)
}
