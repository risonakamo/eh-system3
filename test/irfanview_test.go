package eh_system

import (
	"eh_system/lib/irfanview"
	"testing"
)

func Test_IrfanviewGenerate(t *testing.T) {
    irfanview.GenerateThumbnail(
        "./test-images/webp1.webp",
        "output/webp1.jpg",
        100,
    )
}