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
    Title string `json:"title"`
    // total number of items in the album (recursive)
    Items int `json:"items"`
    // items at the top level of the album
    ImmediateItems int `json:"immediateItems"`

    // path to random img in the album
    Img string `json:"img"`
    // last modified date of the album
    Date string `json:"date"`
    // if this is a leaf album or not (only contains images and no other folders)
    Album bool `json:"album"`
}

/* response to request for single album's contents. contains flat list of all images
   in some order */
type AlbumResponse struct {
    Urls []string `json:"urls"`
    Mode ApiMode `json:"apiMode"`
}

/* program configuration file */
type EhSystemConfig struct {
    ImageDir string `yaml:"imageDir"`
    ThumbnailDir string `yaml:"thumbnailDir"`
}

/* cli args */
type EhSystemArgs struct {
    // name of config file, relative to config folder of this app, without file extension
    ConfigName string
}

// represents a thumbnail that is missing. has path of the original
// item and the thumbnail path (that is not existing and needs to
// be created
type MissingThumbnail struct {
    srcItem string
    neededThumbnail string
}