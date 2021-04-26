package main

import "fmt"

func main() {
  fmt.Println("Start Fizz Buzz game")
  for i := 0; i < 100; i++ {
    fmt.Println(i)
    if i%3 == 0 {
      fmt.Println("Fizz")
    } else if i%5 == 0 {
      fmt.Println("Buzz")
    }
  }
  fmt.Println("Game over")
}