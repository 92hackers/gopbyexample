package main

import (
	fmt "fmt"
	big "math/big"
)

type MyBigInt struct {
	*big.Int
}

func (a MyBigInt) Gop_Add(b MyBigInt) MyBigInt {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/38-Overload-operator/overload_op.gop:12
	return MyBigInt{new(big.Int).Add(a.Int, b.Int)}
}
func (a MyBigInt) Gop_AddAssign(b MyBigInt) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/38-Overload-operator/overload_op.gop:16
	a.Int.Add(a.Int, b.Int)
}
func Int(v *big.Int) MyBigInt {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/38-Overload-operator/overload_op.gop:8
	return MyBigInt{v}
}
func (a MyBigInt) Gop_Neg() MyBigInt {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/38-Overload-operator/overload_op.gop:20
	return MyBigInt{new(big.Int).Neg(a.Int)}
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/38-Overload-operator/overload_op.gop:23
	a := Int(big.NewInt(1))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/38-Overload-operator/overload_op.gop:24
	a.Gop_AddAssign(Int(big.NewInt(2)))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/38-Overload-operator/overload_op.gop:25
	fmt.Println(a.Gop_Add(Int(big.NewInt(3))))
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/38-Overload-operator/overload_op.gop:26
	fmt.Println(a.Gop_Neg())
}
