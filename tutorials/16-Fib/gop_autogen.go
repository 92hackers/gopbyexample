package main

import fmt "fmt"

func fib(n int) int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/16-Fib/fib.gop:4
	if n == 0 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/16-Fib/fib.gop:5
		return 0
	} else
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/16-Fib/fib.gop:6
	if n == 1 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/16-Fib/fib.gop:7
		return 1
	} else {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/16-Fib/fib.gop:9
		return fib(n-1) + fib(n-2)
	}
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/16-Fib/fib.gop:14
	fmt.Println(`fib(27):`, fib(27))
}
