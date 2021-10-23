package main

import fmt "fmt"

type student struct {
	name  string
	score int
}

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:6
	a := []student{student{"du", 84}, student{"wang", 70}, student{"ken", 100}}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:8
	hasFullMark := func() (_gop_ok bool) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:8
		_, x := range a {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:8
			if x.score == 100 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:8
				return true
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:8
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:9
	fmt.Println("is any student full mark:", hasFullMark)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:11
	hasFailed := func() (_gop_ok bool) {
		for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:11
		_, x := range a {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:11
			if x.score < 60 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:11
				return true
			}
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:11
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/13-Exists-comprehension/exists.gop:12
	fmt.Println("is any student failed:", hasFailed)
}
