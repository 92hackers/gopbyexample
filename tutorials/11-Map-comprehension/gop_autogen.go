package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:1
	y := func() (_gop_ret map[int]int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:1
		_gop_ret = map[int]int{}
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:1
		i, x := range []int{1, 3, 5, 7, 11} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:1
			_gop_ret[x] = i
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:1
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:2
	fmt.Println(y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:4
	y = func() (_gop_ret map[int]int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:4
		_gop_ret = map[int]int{}
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:4
		i, x := range []int{1, 3, 5, 7, 11} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:4
			if i%2 == 1 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:4
				_gop_ret[x] = i
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:4
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:5
	fmt.Println(y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:7
	z := func() (_gop_ret map[string]int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:7
		_gop_ret = map[string]int{}
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:7
		k, v := range map[int]string{1: "Hello", 3: "Hi", 5: "xsw", 7: "Go+"} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:7
			if k > 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:7
				_gop_ret[v] = k
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:7
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/11-Map-comprehension/map_comprehens.gop:8
	fmt.Println(z)
}
