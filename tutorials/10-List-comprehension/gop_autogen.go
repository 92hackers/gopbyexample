package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:1
	y := func() (_gop_ret []int) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:1
		_, x := range []int{1, 3, 5, 7, 11} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:1
			_gop_ret = append(_gop_ret, x*x)
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:1
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:2
	fmt.Println(y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:4
	y = func() (_gop_ret []int) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:4
		_, x := range []int{1, 3, 5, 7, 11} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:4
			if x > 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:4
				_gop_ret = append(_gop_ret, x*x)
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:4
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:5
	fmt.Println(y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:7
	z := func() (_gop_ret []int) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:7
		i, v := range []int{1, 3, 5, 7, 11} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:7
			if i%2 == 1 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:7
				_gop_ret = append(_gop_ret, i+v)
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:7
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:8
	fmt.Println(z)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:10
	fmt.Println(func() (_gop_ret []string) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:10
		k, s := range map[string]string{"Hello": "xsw", "Hi": "Go+"} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:10
			_gop_ret = append(_gop_ret, k+","+s)
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:10
		return
	}())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:12
	arr := []int{1, 2, 3, 4, 5, 6}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:13
	x := func() (_gop_ret [][]int) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:13
		_, b := range arr {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:13
			if b > 2 {
				for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:13
				_, a := range arr {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:13
					if a < b {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:13
						_gop_ret = append(_gop_ret, []int{a, b})
					}
				}
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:13
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/10-List-comprehension/list_comprehens.gop:14
	fmt.Println("x:", x)
}
