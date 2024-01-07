package eh_system

import (
	eh_system "eh_system/lib"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// working dir for dumping test outputs
var outputDir string

var testImagesDir string

func TestMain(m *testing.M) {
    // preparing output dump dir
    var cwd string
    cwd,_=os.Getwd()

    outputDir=filepath.Join(cwd,"output")
    testImagesDir=filepath.Join(cwd,"test-images")

    err:=os.RemoveAll(outputDir)
    if err!=nil {
        log.Fatal(err)
    }

    os.Mkdir(outputDir,0755)

    m.Run()
}

// test generate single thumbnail function. dumps into output folder. test images are NOT included.
func TestGenerateThumbnail(t *testing.T) {
    var testImages []string=[]string {
        "tall1.jpg",
        "wide1.jpg",
        "wide2.gif",
        "wide3.mp4",
    }

    for i:=range testImages {
        t.Run(fmt.Sprint(i),func (t *testing.T) {
            var imagePath string=filepath.Join(testImagesDir,testImages[i])

            var err error=eh_system.GenThumbnail(
                imagePath,
                eh_system.ImagePathToThumbnailPath(imagePath,outputDir),
                100,
            )

            if err!=nil {
                t.Error("generate returned error")
            }
        })
    }
}