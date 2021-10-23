package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:1
	fmt.Println("start")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:3
	goto L
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:4
	fmt.Println("before")
L:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:6
	fmt.Println("over")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:7
	i := 0
L2:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:9
	if i < 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:10
		fmt.Println(i)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:11
		i++
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:12
		goto L2
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:14
	fmt.Println("over")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:16
	sum := 0
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:17
	arr := []int{1, 3, 5, 7, 11, 13, 17}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:18
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:18
	i = 0; i < len(arr);
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:18
	i++ {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:19
		if arr[i] < 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:20
			continue
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:22
		if arr[i] > 11 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:23
			break
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:25
		sum += arr[i]
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:27
	fmt.Println("sum(3,5,7,11):", sum == 26, sum)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:28
	sum = 0
L3:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:30
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:30
	i = 0; i < len(arr);
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:30
	i++ {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:31
		if arr[i] < 3 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:32
			continue L3
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:34
		if arr[i] > 11 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:35
			break L3
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:37
		sum += arr[i]
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:39
	fmt.Println("sum(3,5,7,11):", sum == 26, sum)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:41
	z := 3
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:42
	v := "Hello"
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:43
	switch z {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:44
	case 3:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:45
		if v == "Hello" {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:46
			fmt.Println("break")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:47
			break
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:49
		fmt.Println("break fail")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:50
	default:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:51
		fmt.Println(z)
	}
L4:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:54
	switch z {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:55
	case 3:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:56
		if v == "Hello" {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:57
			fmt.Println("break")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:58
			break L4
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:60
		fmt.Println("break fail")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:61
	default:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/21-Break-continue-goto/flow.gop:62
		fmt.Println(z)
	}
}
