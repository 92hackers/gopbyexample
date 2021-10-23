package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:4
	sum := 0
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:5
	_, x := range []int{1, 3, 5, 7, 11, 13, 17} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:5
		if x > 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:6
			sum += x
		}
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:8
	fmt.Println("sum(5,7,11,13,17):", sum)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:10
	fns := make([]func() int, 3)
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:11
	i, x := range []int{3, 15, 777} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:12
		v := x
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:13
		fns[i] = func() int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:14
			return v
		}
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:17
	fmt.Println("values:", fns[0](), fns[1](), fns[2]())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:22
	sum = 0
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:23
	_, x := range []int{1, 3, 5, 7, 11, 13, 17} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:24
		if x > 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:25
			sum += x
		}
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:28
	fmt.Println("sum(5,7,11,13,17):", sum)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:30
	sum = 0
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:31
	i, x := range []int{3, 15, 777} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:32
		v := x
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:33
		fns[i] = func() int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:34
			return v
		}
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:37
	fmt.Println("values:", fns[0](), fns[1](), fns[2]())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:42
	sum = 0
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:43
	x := 0
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:44
	_, x = range []int{1, 3, 5, 7, 11, 13, 17} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:45
		if x > 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:46
			sum += x
		}
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:49
	fmt.Println("x:", x, x == 17)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:50
	fmt.Println("sum(5,7,11,13,17):", sum)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:55
	sum = 0
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:56
	arr := []int{1, 3, 5, 7, 11, 13, 17}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:57
	i := 10
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:58
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:58
	i = 0; i < len(arr);
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:58
	i++ {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:59
		if arr[i] > 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:60
			sum += arr[i]
		}
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/22-For-loop/for.gop:63
	fmt.Println("sum(5,7,11,13,17):", sum)
}
