package main

import "fmt"

func main() {
  fmt.Println("Введите три числа:")
  var first, second, third int
  fmt.Scan(&first)
  fmt.Scan(&second)
  fmt.Scan(&third)

 hasEqual:= false
  if(first == second || first == third || second == third){
    hasEqual = true
  }

  if hasEqual{
    fmt.Print("Два числа или более совпадают")
  }else{
    fmt.Print("Все три числа разные")
  }

}