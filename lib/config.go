package eh_system

import (
	"io/fs"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

/* try to load eh system config from target location */
func LoadConfig(path string) EhSystemConfig {
	var rfile []byte
	var err error
	rfile,err= os.ReadFile(path)

    if err!=nil {
        log.Fatalf("config file read error: %v",err)
        panic("failed to read config")
    }

    var config EhSystemConfig

    if yaml.Unmarshal(rfile,&config)!=nil {
        log.Fatalf("yaml read error: %v",err)
    }

    if !isDir(config.ImageDir) {
        log.Fatalf("provided bad image path %v\n",config.ImageDir)
    }

    if !isDir(config.ThumbnailDir) {
        log.Fatalf("provided bad thumbnail path %v\n",config.ThumbnailDir)
    }

    return config
}

/* check if path is a dir */
func isDir(path string) bool {
    var stat fs.FileInfo
    var err error
    stat,err=os.Stat(path)

    if err!=nil {
        log.Printf("failed to use isdir on: %v",path)
        log.Fatalln(err)
    }

    return stat.IsDir()
}