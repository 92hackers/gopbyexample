package main

import fmt "fmt"

var x = "Hello, world!"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:5
	foo := func(prompt string) (n int, err error) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:6
		n, err = fmt.Println(prompt + x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:7
		return
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:10
	fmt.Println(foo("x: "))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:12
	printf := func(format string, args ...interface {
	}) (n int, err error) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:13
		n, err = fmt.Printf(format, args...) 
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:14
		return
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:17
	bar := func(foo func(string, ...interface {
	}) (int, error)) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:18
		foo("Hello, %v!\n", "Go+")
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:21
	bar(printf)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/05-Closure/closure.gop:22
	fmt.Println(printf("Hello, %v\n", "Go+"))
}
