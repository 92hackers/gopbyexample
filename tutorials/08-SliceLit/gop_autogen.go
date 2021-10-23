package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:1
	x := []float64{1, 3.4}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:2
	fmt.Println("x:", x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:4
	y := []int{1}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:5
	fmt.Println("y:", y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:7
	z := []interface {
	}{1 + 2i, "xsw"}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:8
	fmt.Println("z:", z)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:10
	fmt.Println([]complex128{1, 3.4, 3 + 4i})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:11
	fmt.Println([]complex128{5 + 6i})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:12
	fmt.Println([]interface {
	}{"xsw", 3})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/08-SliceLit/slicelit.gop:14
	fmt.Println("empty slice:", []interface {
	}{})
}
