package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:1
	x := []float64{1, 3.4, 5}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:2
	y := map[string]float64{"Hello": 1, "xsw": 3.4}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:4
	a := [...]float64{1, 3.4, 5}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:5
	b := [...]float64{1, 3: 3.4, 5}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:6
	c := []float64{2: 1.2, 3, 6: 4.5}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:8
	x[1], y["xsw"] = 1.7, 2.8
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:9
	fmt.Println("x:", x, "y:", y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:10
	fmt.Println(`x[1]:`, x[1], `y["xsw"]:`, y["xsw"], `a[1]:`, a[1])
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:12
	i := uint16(4)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:13
	b[uint32(4)], c[i] = 123, 1.7
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:14
	fmt.Println("a:", a, "b:", b, "c:", c)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:16
	arr := [...]float64{1, 2}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:17
	title := "Hello,world!" + "2020-05-27"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/06-String-Map-Array-Slice/datastruct.gop:18
	fmt.Println(title[:len(title)-len("2006-01-02")], len(arr), arr[1:])
}
