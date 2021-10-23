package main

import fmt "fmt"

type foo struct {
}

func (p *foo) Gop_Enum(proc func(key int, val string)) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:5
	proc(3, "Hi")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:6
	proc(7, "Go+")
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:9
	new(foo).Gop_Enum(func(k int, v string) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:10
		fmt.Println(k, v)
	})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:13
	fmt.Println(func() (_gop_ret map[string]int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:13
		_gop_ret = map[string]int{}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:13
		new(foo).Gop_Enum(func(k int, v string) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:13
			_gop_ret[v] = k
		})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/41-UDT-RangeForEach/udt_range.gop:13
		return
	}())
}
