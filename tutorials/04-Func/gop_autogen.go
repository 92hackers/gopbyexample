package main

import (
	fmt "fmt"
	strings "strings"
)

func foo(x string) string {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/04-Func/func.gop:7
	return strings.NewReplacer("?", "!").Replace(x)
}
func printf(format string, args ...interface {
}) (n int, err error) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/04-Func/func.gop:11
	n, err = fmt.Printf(format, args...) 
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/04-Func/func.gop:12
	return
}
func bar(foo func(string, ...interface {
}) (int, error)) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/04-Func/func.gop:16
	foo("Hello, %v!\n", "Go+")
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/04-Func/func.gop:19
	bar(printf)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/04-Func/func.gop:20
	fmt.Println(foo("Hello, world???"))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/04-Func/func.gop:21
	fmt.Println(printf("Hello, %v\n", "Go+"))
}
