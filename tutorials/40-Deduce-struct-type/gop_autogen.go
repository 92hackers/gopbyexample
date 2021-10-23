package main

import fmt "fmt"

type Config struct {
	Dir   string
	Level int
}
type Result struct {
	Text string
}

func foo(conf *Config) *Result {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/40-Deduce-struct-type/deduce.gop:11
	fmt.Println("conf:", *conf)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/40-Deduce-struct-type/deduce.gop:12
	return &Result{Text: "Hello, Go+"}
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/40-Deduce-struct-type/deduce.gop:15
	ret := foo(&Config{Dir: "/foo/bar", Level: 1})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/40-Deduce-struct-type/deduce.gop:16
	fmt.Println(ret)
}
