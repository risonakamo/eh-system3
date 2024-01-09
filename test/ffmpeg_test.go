package eh_system

import (
	eh_system "eh_system/lib"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// working dir for dumping test outputs
var outputDir string
var outputDir2 string

var testImagesDir string

// test images should be placed in test-images folder
var testImages []string=[]string {
    "tall1.jpg",
    "wide1.jpg",
    "wide2.gif",
    "wide3.mp4",
}

var testImagesWithBad []string

func TestMain(m *testing.M) {
    // preparing output dump dir
    var cwd string
    cwd,_=os.Getwd()

    outputDir=filepath.Join(cwd,"output")
    outputDir2=filepath.Join(cwd,"output2")
    testImagesDir=filepath.Join(cwd,"test-images")

    os.RemoveAll(outputDir)
    os.RemoveAll(outputDir2)

    os.Mkdir(outputDir,0755)
    os.Mkdir(outputDir2,0755)



    // converting img paths to actual paths
    for i:=range testImages {
        testImages[i]=filepath.Join(testImagesDir,testImages[i])
    }



    // adding a bad img to test imgs with bad
    testImagesWithBad=append(testImagesWithBad,testImages...)
    testImagesWithBad=append(testImagesWithBad,"bad.png")



    m.Run()
}

// test generate single thumbnail function. dumps into output folder. test images are NOT included.
func TestGenerateThumbnail(t *testing.T) {
    for i:=range testImages {
        t.Run(fmt.Sprint(i),func (t *testing.T) {
            var imagePath string=testImages[i]

            var err error=eh_system.GenThumbnail(
                imagePath,
                eh_system.ImagePathToThumbnailPath(imagePath,outputDir),
                100,
                false,
            )

            if err!=nil {
                t.Error("generate returned error")
            }
        })
    }
}

// test generate parallel on test imgs with a single bad img. it should print out that the
// bad img failed, but the whole thing should not fail.
// also tests with a mismatching array, which should cause error
func TestGenerateParallel(t *testing.T) {
    var outputImgs []string

    t.Run("1",func(t *testing.T) {
        for i := range testImagesWithBad {
            outputImgs=append(outputImgs,eh_system.ImagePathToThumbnailPath(
                testImagesWithBad[i],
                outputDir2,
            ))
        }

        var err error=eh_system.GenThumbnails(
            testImagesWithBad,
            outputImgs,
            100,
            5,
            true,
        )

        if err!=nil {
            t.Error("error from gen thumbnails")
        }
    })

    t.Run("2",func(t *testing.T) {
        for i := range testImagesWithBad {
            outputImgs=append(outputImgs,eh_system.ImagePathToThumbnailPath(
                testImagesWithBad[i],
                outputDir2,
            ))
        }

        var err error=eh_system.GenThumbnails(
            testImages,
            outputImgs,
            100,
            5,
            true,
        )

        if err!=nil {
            return
        }

        t.Error("did not properly return error")
    })
}