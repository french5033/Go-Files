package main

import ("fmt"
        "math"
        "math/rand")

func add(x,y float64) float64{
  return x+y
}

func multiple(a,b string) (string,string) {
    return a,b
}

func main() {
     fmt.Println("The square root of 4 is", math.Sqrt(4))
    fmt.Println("A number from 1-100: ",rand.Intn(100))

    //var num1,num2 float64 = 5.6, 9.5
    num1,num2 := 5.6, 9.5
    fmt.Println(add(num1,num2))

    add(10,20)

    w1,w2 := "Hey","Jude"
    fmt.Println(multiple(w1,w2))

    var a int = 62
    var b float64 = float64(a)
    fmt.Println(a, b)

    x := a
    fmt.Println(x)
}
