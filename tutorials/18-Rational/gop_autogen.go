package main

import (
	fmt "fmt"
	builtin "github.com/goplus/gop/builtin"
	big "math/big"
)

var a builtin.Gop_bigint = builtin.Gop_bigint_Init__1(func() *big.Int {
	v, _ := new(big.Int).SetString("36893488147419103232", 10)
	return v
}())
var b builtin.Gop_bigrat = builtin.Gop_bigrat_Init__2(big.NewRat(4, 5))

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:5
	c := b.Gop_Sub(builtin.Gop_bigrat_Init__2(big.NewRat(1, 3))).Gop_Add(builtin.Gop_bigrat_Init__2(big.NewRat(3, 2)))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:6
	fmt.Println(a, b, c)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:8
	var x *big.Int = func() *big.Int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:8
		v, _ := new(big.Int).SetString("36893488147419103232", 10)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:8
		return v
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:9
	var y *big.Rat = big.NewRat(4, 5)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:10
	fmt.Println(x, y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:12
	a = builtin.Gop_bigint_Init__1(new(big.Int).Abs(big.NewInt(-265)))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/18-Rational/rational.gop:13
	fmt.Println("abs(-265r):", a)
}
