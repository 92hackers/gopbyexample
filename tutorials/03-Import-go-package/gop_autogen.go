package main

import (
	fmt "fmt"
	strings "strings"
)

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/03-Import-go-package/import.gop:6
	x := strings.NewReplacer("?", "!").Replace("hello, world???")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/03-Import-go-package/import.gop:7
	fmt.Println("x:", x)
}
