package main

import fmt "fmt"

func A(a *int, c *struct {
	b *int
	m map[string]*int
	s []*int
}) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:6
	*a = 5
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:7
	*c.b = 3
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:8
	*c.m["foo"] = 7
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:9
	*c.s[0] = 9
}
func Index() int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:13
	return 0
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:16
	a1 := 6
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:17
	a2 := 6
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:18
	a3 := 6
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:19
	c := struct {
		b *int
		m map[string]*int
		s []*int
	}{b: &a1, m: map[string]*int{"foo": &a2}, s: []*int{&a3}}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:30
	A(&a1, &c)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:31
	*c.m["foo"] = 8
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:32
	*c.s[0] = 10
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:33
	*c.s[Index()] = 11
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/27-Func-Set/func.gop:34
	fmt.Println(a1, *c.b, *c.m["foo"], *c.s[0])
}
