package main

import fmt "fmt"

func foo() []int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:2
	return nil
}
func foo1() map[int]int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:6
	return make(map[int]int, 10)
}
func foo2() chan int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:10
	return make(chan int, 10)
}
func foo3() *int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:14
	return nil
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:17
	fmt.Println(foo() == nil)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:18
	fmt.Println(nil == foo())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:19
	fmt.Println(foo() != nil)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:20
	fmt.Println(nil != foo())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:22
	fmt.Println(foo1() == nil)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:23
	fmt.Println(nil == foo1())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:24
	fmt.Println(foo1() != nil)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:25
	fmt.Println(nil != foo1())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:27
	fmt.Println(foo2() == nil)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:28
	fmt.Println(nil == foo2())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:29
	fmt.Println(foo2() != nil)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:30
	fmt.Println(nil != foo2())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:32
	fmt.Println(foo3() == nil)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:33
	fmt.Println(nil == foo3())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:34
	fmt.Println(foo3() != nil)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/29-CompareToNil/ref.gop:35
	fmt.Println(nil != foo3())
}
