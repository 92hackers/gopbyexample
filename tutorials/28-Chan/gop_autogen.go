package main

import fmt "fmt"

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/28-Chan/chan.gop:1
	c := make(chan int, 10)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/28-Chan/chan.gop:2
	c <- 3
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/28-Chan/chan.gop:3
	close(c)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/28-Chan/chan.gop:5
	d := <-c
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/28-Chan/chan.gop:6
	e, ok := <-c
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/28-Chan/chan.gop:8
	fmt.Println(d)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/28-Chan/chan.gop:9
	fmt.Println(e, ok)
}
