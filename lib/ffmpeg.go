package eh_system

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// generate thumbnail of target file. places at target output.
// output needs to have extension (probably .jpg). thumbnail
// will always be square with the specified dimensions
func GenThumbnail(
	targetFile string,
	outputFile string,
    size int,
    suppressFfmpegError bool,
) error {
    var err error=os.MkdirAll(filepath.Dir(outputFile),0655)

    if err!=nil {
        color.Red("error while trying to make directory for:",outputFile)
        return err
    }

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
    out,err=cmd.CombinedOutput()

    if err!=nil {
        color.Red("error while generating thumbnail: %v",targetFile)

        if !suppressFfmpegError {
            color.Red("ffmpeg error text:")
            fmt.Println(string(out))
        }

        return err
    }

    fmt.Fprintf(color.Output,"generated thumbnail: %v\n",color.YellowString(outputFile))
    return nil
}

// use gen thumbnails to gen thumbnails in parallel.
// failed generations do not cause this function
// to end in error, errors will only be reported.
//
// provide list of input files, and list of output files. output files must have file extension
// (probably .jpg). errors if size of inputs does not match size of outputs.
func GenThumbnails(
    inputs []string,
    outputs []string,
    size int,
    workers int,
    suppressFfmpegError bool,
) error {
    if len(inputs)!=len(outputs) {
        color.Red("got differing sized input/output arrays. "+
            "input/output arrays must be the same size")
        fmt.Println("input size:",len(inputs))
        fmt.Println("output size:",len(outputs))
        return errors.New("input/output size difference")
    }

    var thumbnailJobs []ThumbnailJob

    // building thumbnail jobs array
    for i := range inputs {
        thumbnailJobs=append(thumbnailJobs,ThumbnailJob{
            input:inputs[i],
            output:outputs[i],
            size:size,
        })
    }

    var wg sync.WaitGroup
    var jobsChannel chan ThumbnailJob=make(chan ThumbnailJob)

    // spawn workers
    for i:=0;i<workers;i++ {
        wg.Add(1)
        go thumbnailGenWorker(jobsChannel,size,&wg,suppressFfmpegError)
    }

    // submit jobs to workers
    for i := range thumbnailJobs {
        jobsChannel<-thumbnailJobs[i]
    }

    close(jobsChannel)
    wg.Wait()

    return nil
}

// worker version of thumbnail gen. works on jobs recved on job channel until it closes.
// if a job fails, does not die, just keeps going.
func thumbnailGenWorker(
    jobChannel <-chan ThumbnailJob,
    size int,
    wg *sync.WaitGroup,
    suppressFfmpegError bool,
) {
    // work until the channel closes
    for {
        var job ThumbnailJob
        var ok bool
        job,ok=<-jobChannel

        if !ok {
            break
        }

        GenThumbnail(job.input,job.output,size,suppressFfmpegError)
    }

    wg.Done()
}

// convert a path to an image to a path to a thumbnail of the image, if it were to be placed in
// the target dir. gives jpg extension.
func ImagePathToThumbnailPath(
    srcImg string,
    outputDir string,
) string {
    var filename string=strings.TrimSuffix(
        filepath.Base(srcImg),
        filepath.Ext(srcImg),
    )

    return filepath.Join(outputDir,filename+".jpg")
}