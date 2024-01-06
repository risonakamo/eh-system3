package main

import "github.com/disintegration/gift"

func main() {
	g := gift.New(
        gift.CropToSize(100,100,gift.CenterAnchor),
    )

    // g.Bounds()
}