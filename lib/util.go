// random utils

package eh_system

import (
	"io/fs"
	"math/rand"
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