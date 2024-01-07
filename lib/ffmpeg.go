package eh_system

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

/* generate thumbnail of target file. places at target output.
   output needs to have extension (probably .jpg). thumbnail
   will always be square with the specified dimensions */
func GenThumbnail(
	targetFile string,
	outputFile string,
    size int,
) error {
	var cmd *exec.Cmd=exec.Command(
        "ffmpeg.exe",
        "-i",targetFile,
        "-vf",
        "thumbnail,"+
            fmt.Sprintf(
                "scale='if(gt(iw,ih),-2,%d)':'if(gt(iw,ih),%d,-2)',",
                size,size,
            )+
            fmt.Sprintf("crop=%d:%d",size,size),
        "-frames:v","1",
        outputFile,
    )

    var out []byte
    var err error
    out,err=cmd.CombinedOutput()

    if err!=nil {
        color.Red("error while generating thumbnail: %v",targetFile)
        color.Red("ffmpeg error text:")
        fmt.Println(string(out))
        return err
    }

    fmt.Fprintf(color.Output,"generated thumbnail: %v",color.YellowString(outputFile))
    return nil
}