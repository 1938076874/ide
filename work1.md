package main



import (

  "fmt"

  "time"

)



func name(ch chan string) {

  for i := 1; i <= 10; i++ {

​    ch <- "张三"

​    ch <- "李四"

​    ch <- "王五"

​    ch <- "赵六"

​    a := <-ch

​    b := <-ch

​    c := <-ch

​    d := <-ch

​    fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)

  }

}



func main() {

  intChan := make(chan string, 10)

  go name(intChan)

  time.Sleep(10 * time.Second)

}