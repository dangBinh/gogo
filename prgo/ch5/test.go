package main

import (
  "fmt"
)

func main() {
  var i interface{} = 99
  var s interface{} = []string{"left", "right"}
  j := i.(int)
  fmt.Printf("%T->%d\n", j, j)
  if i, ok := i.(int); ok { // shadow variable of int
    fmt.Printf("%T->%d\n", i, i)
  }
  if s, ok := s.([]string); ok {
    fmt.Printf("%T->%q\n", s, s)
  }

  counterA := createCounter(2)
  counterB := createCounter(102)
  for countVar := 1; countVar <= 5; countVar ++ {
    a := <- counterA
    fmt.Printf("(A->%d, B->%d)  ", a, <-counterB)
  }

  // Generic function
  iType := Minimum(4, 3, 2, 8, 9).(int)
  fmt.Printf("%T %v\n", iType, iType)
  fType := Minimum(9.4, 9.5, 9.6).(float64)
  fmt.Printf("%T %v\n", fType, fType)
  sType := Minimum("A", "B", "C").(string)
  fmt.Printf("%T %v\n", sType, sType)
}

func createCounter(start int) chan int {
  next := make(chan int) // no capacity
  go func(i int) {
    for {
      next <- i
      i++
    }
  }(start)
  return next
}

func Minimum(first interface{}, rest ...interface{}) interface{} {
  minimum := first
  for _, x := range rest {
    switch x := x.(type) {
    case int:
      if x < minimum.(int) {
        minimum = x
      }
    case float64:
      if x < minimum.(float64) {
        minimum = x
      }
    case string:
      if x < minimum.(string) {
        minimum = x
      }
    }
  }
  return minimum
}
