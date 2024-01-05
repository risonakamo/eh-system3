package eh_system

import (
	"log"
	"os"

	"github.com/akamensky/argparse"
)

/* get eh system cli args */
func GetArgs() EhSystemArgs {
	var parser *argparse.Parser= argparse.NewParser(
        "eh-system v4",
        "eh system server",
    )

    var configFile *string=parser.String(
        "c",
        "config",
        &argparse.Options{
            Required:true,
            Help:"config file name, relative to config folder of this app. no file extension",
        },
    )

    var err error=parser.Parse(os.Args)
    if err!=nil {
        log.Fatalln(err)
    }

    return EhSystemArgs {
        ConfigName:*configFile,
    }
}