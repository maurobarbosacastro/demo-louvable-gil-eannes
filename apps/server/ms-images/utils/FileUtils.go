package utils

import (
    "image"
)

const (
    Horizontal string = "horizontal"
    Vertical          = "vertical"
    Square            = "square"
)

func GetOrientation(imageFile image.Image) string {
    if imageFile.Bounds().Dx() > imageFile.Bounds().Dy() {
        return Horizontal
    } else if imageFile.Bounds().Dx() < imageFile.Bounds().Dy() {
        return Vertical
    }
    return Square
}

func IsHorizontal(orientation string) bool {
    return orientation == Horizontal
}
func IsVertical(orientation string) bool {
    return orientation == Vertical
}
func IsSquare(orientation string) bool {
    return orientation == Square
}
