package eh_system

import (
	"eh_system/lib/irfanview"
	"testing"
)

func Test_IrfanviewGenerate(t *testing.T) {
    irfanview.GenerateThumbnail(
        "./test-images/webp3.webp",
        "output/webp3.jpg",
        100,
    )
}