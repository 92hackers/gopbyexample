package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/07-MapLit/maplit.gop:1
	x := map[string]float64{"Hello": 1, "xsw": 3.4}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/07-MapLit/maplit.gop:2
	fmt.Println("x:", x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/07-MapLit/maplit.gop:4
	y := map[string]interface {
	}{"Hello": 1, "xsw": "Go+"}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/07-MapLit/maplit.gop:5
	fmt.Println("y:", y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/07-MapLit/maplit.gop:7
	fmt.Println(map[string]int{"Hello": 1, "xsw": 3})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/07-MapLit/maplit.gop:8
	fmt.Println(map[int]interface {
	}{1: 1.4, 3: "Go+"})
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/07-MapLit/maplit.gop:10
	fmt.Println("empty map:", map[string]interface {
	}{})
}
