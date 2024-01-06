package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bakape/thumbnailer/v2"
)

func main() {
	file,_:=os.Open("test.png")

    _,_,err:=thumbnailer.Process(file,thumbnailer.Options{
        ThumbDims: thumbnailer.Dims{
            Width: 150,
            Height: 150,
        },
    })

    if err!=nil {
        fmt.Println("died")
        log.Fatalln(err)
    }

    // reader:=io.ReadSeeker(file)

    // _,img,_:=thumbnailer.Process(reader,thumbnailer.Options{})

    // outfile,_:=os.Create("a.jpg")
    // jpeg.Encode(outfile,img,nil)

}