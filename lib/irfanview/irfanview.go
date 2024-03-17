package irfanview

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// generate a thumbnail using irfanview
// output file must include extension
func GenerateThumbnail(
	targetFile string,
	outputFile string,
	size int,
) {
    var e error
    targetFile,e=filepath.Abs(targetFile)

    if e!=nil {
        panic(e)
    }

    outputFile,e=filepath.Abs(outputFile)

    if e!=nil {
        panic(e)
    }

	e=os.MkdirAll(filepath.Dir(outputFile), 0655)

    if e!=nil {
        fmt.Println("failed to make dir")
        panic(e)
    }

    var cmd *exec.Cmd=exec.Command(
        "i_view64.exe",
        targetFile,
        fmt.Sprintf(
            "/resize_short=%d",
            size,
        ),
        "/aspectratio",
        "/resample",
        fmt.Sprintf(
            "/crop=(0,0,%d,%d)",
            size,
            size,
        ),
        fmt.Sprintf("/convert=%s",outputFile),
    )

    var out []byte
    out,e=cmd.CombinedOutput()

    if e!=nil {
        fmt.Printf("error while irfanview generating thumbnail: %v\n",targetFile)

        fmt.Println("error text:")
        fmt.Println(string(out))

        panic(e)
    }

    fmt.Printf("generated thumbnail: %s\n",outputFile)
}