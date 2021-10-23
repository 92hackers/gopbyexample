package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:1
	x := 0
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:2
	if
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:2
	t := false; t {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:3
		x = 3
	} else {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:5
		x = 5
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:7
	fmt.Println("x:", x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:9
	x = 0
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:10
	switch
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:10
	s := "Hello"; s {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:11
	default:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:12
		x = 7
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:13
	case "world", "hi":
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:14
		x = 5
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:15
	case "xsw":
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:16
		x = 3
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:18
	fmt.Println("x:", x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:20
	v := "Hello"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:21
	switch {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:22
	case v == "xsw":
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:23
		x = 3
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:24
	case v == "Hello", v == "world":
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:25
		x = 9
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:26
	default:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:27
		x = 7
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:29
	fmt.Println("x:", x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:31
	v = "Hello"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:32
	switch {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:33
	case v == "xsw":
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:34
		x = 3
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:35
	case v == "hi", v == "world":
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:36
		x = 9
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:37
	default:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:38
		x = 11
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:40
	fmt.Println("x:", x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:42
	switch v {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:43
	case "Hello":
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:44
		fmt.Println(v)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:44
		fallthrough
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:46
	case "hi":
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:47
		fmt.Println(v)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:47
		fallthrough
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:49
	default:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:50
		fmt.Println(v)
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:53
	z := 3
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:54
	switch {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:55
	case z < 10:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:56
		fmt.Println(z)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:56
		fallthrough
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:58
	case z == 10:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:59
		fmt.Println(z)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:59
		fallthrough
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:61
	case z > 10:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:62
		fmt.Println(z)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:62
		fallthrough
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:64
	default:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/09-IfElse-SwitchCase/flow.gop:65
		fmt.Println(z)
	}
}
