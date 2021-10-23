package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:1
	n := 2
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:2
	a := make([]int, uint64(n))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:3
	a = append(a, 1, 2, 3)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:4
	fmt.Println(a)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:5
	fmt.Println("len:", len(a))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:7
	b := make([]int, 0, uint16(4))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:8
	c := []int{1, 2, 3}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:9
	b = append(b, c...) 
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:10
	fmt.Println(b)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/31-Builtin-Typecast/builtin_and_typecast.gop:11
	fmt.Println("len:", len(b), "cap:", cap(b))
}
