package main

import (
	"bytes"
	"net/http"
)

func main() {
    var r *http.Request
    r,_=http.NewRequest(
        "POST",
        "http://localhost:4200/get-album",
        bytes.NewBufferString("nekonote"),
    )

    var client *http.Client=&http.Client{}
    client.Do(r)
}