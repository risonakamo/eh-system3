// random utils

package eh_system

import (
	"io/fs"
	"math/rand"
	"path/filepath"
	"strings"
)

/** shuffle an array */
func shuffleArray[T any](array []T) {
    rand.Shuffle(len(array),func (i int,j int) {
        (array)[i],(array)[j]=(array)[j],(array)[i]
    })
}

/** checks if all items in fs dir entry array is a file */
func isAllFiles(items []fs.DirEntry) bool {
    for i := range items {
        if items[i].IsDir() {
            return false
        }
    }

    return true
}

/** random pick from array */
func randFromArray[T any](array []T) T {
    return array[rand.Intn(len(array))]
}

// replace extension of a file path with a new one. new ext should not have
// a dot.
// does not matter if it is a partial path or full path, just that the path
// ends with an extension
func replaceExt(path string,newExt string) string {
    var basename string=strings.TrimSuffix(
        path,
        filepath.Ext(path),
    )

    return basename+"."+newExt
}