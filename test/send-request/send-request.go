package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
    var req *http.Request
    req,_=http.NewRequest(
        "POST",
        "http://localhost:4200/get-album-info",
        bytes.NewBufferString(""),
    )

    var client *http.Client=&http.Client{}

    var res *http.Response
    res,_=client.Do(req)

    body,_:=io.ReadAll(res.Body)
    fmt.Println(string(body))
}