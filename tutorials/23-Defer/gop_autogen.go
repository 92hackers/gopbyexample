package main

import fmt "fmt"

func f() (x int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:2
	defer func() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:3
		x = 3
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:5
	return 1
}
func g() (x int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:9
	defer func() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:10
		x = 3
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:12
	x = 1
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:13
	return
}
func h() (x int) {
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:17
	_, i := range []int{3, 2, 1} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:18
		v := i
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:19
		defer func() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:20
			x = v
		}()
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:23
	return
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:26
	fmt.Println("f-x:", f())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:27
	fmt.Println("g-x:", g())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:28
	fmt.Println("h-x:", h())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:30
	defer fmt.Println(fmt.Println("Hello, defer"))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/23-Defer/defer.gop:31
	fmt.Println("Hello, Go+")
}
