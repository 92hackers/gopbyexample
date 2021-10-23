package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:1
	a := []int{2, 3, 5, 7, 9, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:2
	where := func() (_gop_ret int) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:2
		i, v := range a {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:2
			if v == 19 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:2
				return i
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:2
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:3
	fmt.Println("19 is at index", where)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:5
	if
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:5
	at, ok := func() (_gop_ret int, _gop_ok bool) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:5
		i, v := range a {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:5
			if v == 37 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:5
				return i, true
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:5
		return
	}(); ok {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:6
		fmt.Println("37 is found! index =", at)
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:9
	fmt.Println("first element of a is:", func() (_gop_ret int) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:9
		_, v := range a {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:9
			return v
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension/select.gop:9
		return
	}())
}
