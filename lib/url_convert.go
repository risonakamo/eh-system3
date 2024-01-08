// functions for dealing with conversion between relative paths and web urls

package eh_system

import (
	"path/filepath"
	"strings"
)

/* change image urls in album infos to be a full web url to thumbnail dir. MUTATES original array */
func FixAlbumInfoImageUrls(albumInfos []AlbumInfo) []AlbumInfo {
    for i := range albumInfos {
        albumInfos[i].Img="/thumbnaildata/"+
            strings.TrimSuffix(
                albumInfos[i].Img,
                filepath.Ext(albumInfos[i].Img),
            )+
            ".jpg"
    }

    return albumInfos
}

/* modify array of local image urls to be web url address. MUTATES original array */
func FixImageUrls(urls []string) []string {
    for i := range urls {
        urls[i]="/imagedata/"+urls[i]
    }

    return urls
}

// given a path to an image **relative to the imageDir**, return the corresponding thumbnail
// path for that image. the thumbnail's extension is always jpg, and it will be relative to
// the thumbnails dir instead
func findThumbnailOfImage(
    imageDir string,
    thumbnailDir string,
    image string,
) string {
    var relPath string
    relPath,_=filepath.Rel(imageDir,image)
    return filepath.Join(thumbnailDir,relPath)
}