package morestrings

import (
	"os"
)

func Create(){
    directories := [3]string{"mypackage", "cmd", "std"}
    for _, dir := range directories{
        os.Mkdir("./" + string(dir), 0755)
    }
}