package main

import fmt "fmt"

func Map(c []float64, t func(float64) float64) []float64 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/39-Lambda-expression/lambda.gop:2
	return func() (_gop_ret []float64) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/39-Lambda-expression/lambda.gop:2
		_, x := range c {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/39-Lambda-expression/lambda.gop:2
			_gop_ret = append(_gop_ret, t(x))
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/39-Lambda-expression/lambda.gop:2
		return
	}()
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/39-Lambda-expression/lambda.gop:5
	fmt.Println(Map([]float64{1.2, 3.5, 6}, func(x float64) float64 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/39-Lambda-expression/lambda.gop:5
		return x * x
	}))
}
