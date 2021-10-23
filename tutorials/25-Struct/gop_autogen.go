package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:1
	a := struct {
		A int
		B string
	}{1, "Hello"}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:6
	fmt.Println(a)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:8
	b := &struct {
		A int
		B string
	}{1, "Hello"}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:13
	c := &struct {
		a int
		b string
	}{1, "Hello"}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:18
	fmt.Println(b)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:20
	a.A, a.B = 1, "Hi"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:21
	fmt.Println(a)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:23
	b.A, b.B = 2, "Hi2"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:24
	fmt.Println(b)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:26
	c.a, c.b = 3, "Hi3"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/25-Struct/struct.gop:27
	fmt.Println(c)
}
