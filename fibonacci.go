package main
import "fmt"

func fibonacci() func() int {
  a, b := 0, 1
    return func() int{
      a,b = b, a+b
        return a
      }
}

func main() {
  f := fibonacci()
  fmt.Println(f(),f(),f(),f(),f(),f(),f(),f(),f())
}
