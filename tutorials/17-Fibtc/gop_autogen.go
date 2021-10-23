package main

import fmt "fmt"

func fibtc(n int, a int, b int) int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/17-Fibtc/fibtc.gop:4
	if n == 0 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/17-Fibtc/fibtc.gop:5
		return a
	} else
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/17-Fibtc/fibtc.gop:6
	if n == 1 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/17-Fibtc/fibtc.gop:7
		return b
	} else {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/17-Fibtc/fibtc.gop:9
		return fibtc(n-1, b, a+b)
	}
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/17-Fibtc/fibtc.gop:14
	fmt.Println(`fibtc(27):`, fibtc(27, 0, 1))
}
