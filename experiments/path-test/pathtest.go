package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	var topdir string = "C:\\Users\\ktkm\\Desktop\\h\\cg"
	var path string = "C:\\Users\\ktkm\\Desktop\\h\\cg\\leonat\\hoshino\\0.jpg"

	fmt.Println(filepath.Rel(topdir,path))
}