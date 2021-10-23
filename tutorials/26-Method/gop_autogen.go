package main

import fmt "fmt"

type Person struct {
	Name    string
	Age     int
	Friends []string
}

func (p *Person) SetName(name string) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:8
	p.Name = name
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:9
	fmt.Println(p.Name)
}
func (p *Person) SetAge(age int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:13
	age = age + 5
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:14
	p.Age = age
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:15
	fmt.Println(p.Age)
}
func (p *Person) AddFriends(args ...string) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:19
	p.Friends = append(p.Friends, args...)
}

type M int

func (m M) Foo() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:25
	fmt.Println("foo", m)
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:28
	p := Person{Name: "bar", Age: 30}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:33
	p.Name, p.Age = "bar2", 31
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:35
	p.SetName("foo")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:36
	p.SetAge(32)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:37
	p.AddFriends("a", "b", "c")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:39
	a := int32(0)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:40
	m := M(a)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:41
	m.Foo()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:43
	fmt.Println(p.Name)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:44
	fmt.Println(p.Age)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:45
	fmt.Println(p.Friends)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/26-Method/method.gop:46
	fmt.Println(m)
}
