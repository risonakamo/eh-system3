package main2

import (
	"log"
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
    rfile,_:=os.Open("test.yml")
	decoder:=yaml.NewDecoder(rfile)

    decoder.KnownFields(true)

    var obj TestYaml

    err:=decoder.Decode(&obj)

    if err!=nil {
        log.Fatalf("%v",err)
        panic("??")
    }

    spew.Dump(obj)
}