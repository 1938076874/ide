package main

import (
	"fmt"
	"runtime"
	"sync"
)

func name1(ch chan string) {
	defer wg.Done()
	ch <- "张三"
	ch <- "李四"
	ch <- "王五"
	ch <- "赵六"
	a := <-ch
	b := <-ch
	c := <-ch
	d := <-ch
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}
func name2(ch chan string) {
	defer wg.Done()
	ch <- "张三"
	ch <- "李四"
	ch <- "王五"
	ch <- "赵六"
	a := <-ch
	b := <-ch
	c := <-ch
	d := <-ch
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}
func name3(ch chan string) {
	defer wg.Done()
	ch <- "张三"
	ch <- "李四"
	ch <- "王五"
	ch <- "赵六"
	a := <-ch
	b := <-ch
	c := <-ch
	d := <-ch
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}
func name4(ch chan string) {
	defer wg.Done()
	ch <- "张三"
	ch <- "李四"
	ch <- "王五"
	ch <- "赵六"
	a := <-ch
	b := <-ch
	c := <-ch
	d := <-ch
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(4)

	A := make(chan string, 20)
	B := make(chan string, 20)
	C := make(chan string, 20)
	D := make(chan string, 20)
	for i := 0; i <= 10; i++ {
		go name1(A)
		go name2(B)
		go name3(C)
		go name4(D)
	}
	wg.Wait()
}
