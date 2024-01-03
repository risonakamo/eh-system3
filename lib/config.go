package eh_system

import (
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
        panic("yaml read err")
    }

    return config
}