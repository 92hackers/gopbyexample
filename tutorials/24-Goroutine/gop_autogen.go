package main

import (
	fmt "fmt"
	time "time"
)

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/24-Goroutine/goroutine.gop:3
	go func() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/24-Goroutine/goroutine.gop:4
		fmt.Println("Hello, goroutine!")
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/24-Goroutine/goroutine.gop:7
	time.Sleep(1e8)
}
