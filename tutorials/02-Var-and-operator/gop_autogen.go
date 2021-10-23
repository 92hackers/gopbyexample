package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/02-Var-and-operator/var_and_op.gop:1
	x := 123.1 - 3i
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/02-Var-and-operator/var_and_op.gop:2
	y, z := "Hello, ", 123
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/02-Var-and-operator/var_and_op.gop:3
	fmt.Println(y+"complex:", x+1, "int:", z)
}
