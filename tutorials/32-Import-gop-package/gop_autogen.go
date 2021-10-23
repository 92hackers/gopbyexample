package main

import (
	fmt "fmt"
	foo "github.com/goplus/gop/tutorial/14-Using-goplus-in-Go/foo"
)

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/32-Import-gop-package/import_gop_pkg.gop:3
	rmap := foo.ReverseMap(map[string]int{"Hi": 1, "Hello": 2})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/32-Import-gop-package/import_gop_pkg.gop:4
	fmt.Println(rmap)
}
