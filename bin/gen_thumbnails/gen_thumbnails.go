// thumbnail detection and generation binary.
// reads config in same manner as eh server to get img and thumbnail dir.
// detects all imgs in the img dir that do not have a corresponding thumbnail
// in the thumbnail dir. generates for all those imgs and places in thumbnail dir.

package main

import (
	eh_system "eh_system/lib"
	"os"
	"path/filepath"
)

// square thumbnail size
const thumbnailSize int=160

const thumbnailWorkers int=8

const suppressErrors bool=false

func main() {
    var HERE string
    HERE,_=os.Executable()
    HERE=filepath.Dir(HERE)

    // --- config get ---
    var args eh_system.EhSystemArgs=eh_system.GetArgs()

    var config eh_system.EhSystemConfig=eh_system.LoadConfig(
        filepath.Join(HERE,"../config",args.ConfigName+".yml"),
    )

    eh_system.GenerateMissingThumbnails(
        config.ImageDir,
        config.ThumbnailDir,
        thumbnailSize,

        thumbnailWorkers,
        suppressErrors,
    )
}