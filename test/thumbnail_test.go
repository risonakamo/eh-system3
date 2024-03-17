package eh_system

import (
	eh_system "eh_system/lib"
	"fmt"
	"testing"
)

// test running find missing thumbnails
func TestFindMissingThumbnails(t *testing.T) {
    var res []eh_system.MissingThumbnail=eh_system.FindMissingImagesWithoutThumbnails(
        "C:/Users/ktkm/Desktop/h/cg",
        "C:/Users/ktkm/Desktop/h/thumbnaildatas",
        10,
    )

    fmt.Println(len(res))
    // spew.Dump(res)
}

func TestGenMissingThumbnails(t *testing.T) {
    eh_system.GenerateMissingThumbnails(
        "test-images",
        "output3",
        160,

        5,
        false,
    )
}