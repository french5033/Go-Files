// intro2.go
package main

import "fmt"

func main() {
  x := 15
  a := &x // memory address
  fmt.Println(a)  //0x.....
  fmt.Println(*a) //15
  *a = 5
  fmt.Println(x) //5
  *a = *a**a
  fmt.Println(x)
  fmt.Println(*a)
}
