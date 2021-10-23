package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/30-Recover/recover.gop:1
	defer func() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/30-Recover/recover.gop:2
		var err interface {
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/30-Recover/recover.gop:3
		if
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/30-Recover/recover.gop:3
		err = recover(); err != nil {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/30-Recover/recover.gop:4
			fmt.Println(err)
		}
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/30-Recover/recover.gop:8
	panic("hello recover")
}
