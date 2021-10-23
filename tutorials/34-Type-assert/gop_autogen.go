package main

import fmt "fmt"

func foo(v interface {
}) string {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:2
	if
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:2
	_, ok := v.(bool); ok {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:3
		return "bool"
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:5
	switch v.(type) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:6
	case int:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:7
		return "int"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:8
	case string:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:9
		return "string"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:10
	default:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:11
		return "unknown"
	}
}
func add(v interface {
}, delta interface {
}) interface {
} {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:16
	switch a := v.(type) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:17
	case int:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:18
		return a + delta.(int)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:19
	case float64:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:20
		return a + delta.(float64)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:21
	case string:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:22
		return a + delta.(string)
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:24
	return nil
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:27
	fmt.Println(foo(1), foo("Hi"))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/34-Type-assert/type_assert.gop:28
	fmt.Println(add(4, 3), add("n", "iu"))
}
