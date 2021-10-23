package main

import fmt "fmt"

type fooIter struct {
	data *foo
	idx  int
}
type foo struct {
	key []int
	val []string
}

func (p *fooIter) Next() (key int, val string, ok bool) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:7
	if p.idx < len(p.data.key) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:8
		key, val, ok = p.data.key[p.idx], p.data.val[p.idx], true
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:9
		p.idx++
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:11
	return
}
func (p *foo) Gop_Enum() *fooIter {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:24
	return &fooIter{data: p}
}
func newFoo() *foo {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:20
	return &foo{key: []int{3, 7}, val: []string{"Hi", "Go+"}}
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:27
	obj := newFoo()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:28
	for _gop_it := obj.Gop_Enum(); ; {
		var _gop_ok bool
		k, v, _gop_ok := _gop_it.Next()
		if !_gop_ok {
			break
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:29
		fmt.Println(k, v)
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:32
	fmt.Println(func() (_gop_ret map[string]int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:32
		_gop_ret = map[string]int{}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:32
		for _gop_it := obj.Gop_Enum(); ; {
			var _gop_ok bool
			k, v, _gop_ok := _gop_it.Next()
			if !_gop_ok {
				break
			}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:32
			_gop_ret[v] = k
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/42-UDT-RangeIterator/udt_range_iter.gop:32
		return
	}())
}
