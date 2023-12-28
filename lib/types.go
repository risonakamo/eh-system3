// types

package eh_system

// asdasd

type ApiMode string

/* possible api modes */
const (
    CLOUD_MODE ApiMode="cloud"
    LOCAL_MODE ApiMode="local"
)

/* a top level album item. an album can recursively contain other albums, or be a leaf album, only
   containing images. both are represented by this struct*/
type AlbumInfo struct {
    Title string
    // total number of items in the album (recursive)
    Items int
    // items at the top level of the album
    ImmediateItems int

    // path to random img in the album
    Img string
    // last modified date of the album
    Date string
    // if this is a leaf album or not (only contains images and no other folders)
    Album bool
}

/* response to request for single album's contents. contains flat list of all images
   in some order */
type AlbumResponse struct {
    Urls []string
    Mode ApiMode
}