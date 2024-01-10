package eh_system

import (
	eh_system "eh_system/lib"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestFindMissingThumbnails(t *testing.T) {
    res:=eh_system.FindMissingImagesWithoutThumbnails(
        "C:/Users/ktkm/Desktop/h/cg",
        "C:/Users/ktkm/Desktop/eh-system3/thumbnaildatas/thumbnaildata",
    )

    spew.Dump(res)
}