package main2

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

type TestYaml struct {
	A string
	B string
	C []string
}

func main2() {
    var rfile []byte
    var err error
	rfile,err=os.ReadFile("test.yml")

    if err!=nil {
        panic(err)
    }

    var obj TestYaml

    err=yaml.Unmarshal(rfile,&obj)

    if err!=nil {
        fmt.Println("uh oh")
        panic(err)
    }

    spew.Dump(obj)
}