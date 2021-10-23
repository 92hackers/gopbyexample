package main

import fmt "fmt"

type student struct {
	name  string
	score int
}

func findScore(a []student, name string) (score int, ok bool) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension2/findscore.gop:7
	return func() (_gop_ret int, _gop_ok bool) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension2/findscore.gop:7
		_, x := range a {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension2/findscore.gop:7
			if x.name == name {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension2/findscore.gop:7
				return x.score, true
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension2/findscore.gop:7
		return
	}()
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension2/findscore.gop:10
	a := []student{student{"ken", 95}, student{"john", 90}, student{"tom", 58}}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/12-Select-comprehension2/findscore.gop:11
	fmt.Println(findScore(a, "tom"))
}
