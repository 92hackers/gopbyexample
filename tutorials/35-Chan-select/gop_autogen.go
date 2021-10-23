package main

import fmt "fmt"

var done = make(chan bool, 1)
var exited = make(chan bool, 1)

func consume(xchg chan int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:5
	for {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:6
		select {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:7
		case
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:7
		c := <-xchg:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:8
			fmt.Println(c)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:9
		case
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:9
		<-done:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:10
			fmt.Println("done!")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:11
			exited <- true
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:12
			return
		}
	}
}
func product(xchg chan int, from chan int) {
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:18
	x := range from {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:19
		xchg <- x
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:21
	done <- true
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:24
	from := make(chan int, 10)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:25
	xchg := make(chan int, 1)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:26
	go consume(xchg)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:27
	go product(xchg, from)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:28
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:28
	i := 1; i <= 10;
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:28
	i++ {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:29
		from <- i
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:31
	close(from)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/35-Chan-select/select.gop:32
	<-exited
}
