package main

	
import (
    "fmt"
)

func normalFunction(message string){
  fmt.Println(message)
}

func duplicado( a int) int{
  return a*2
}

func returntwovalue( a int) (c,b int){
  return a,a*2
}


func main() {
  normalFunction("Hola mundo")
  normalFunction(fmt.Sprint("Value: ",duplicado(4)))
  value1, _ := returntwovalue(2)
  normalFunction(fmt.Sprint("Value1: ",value1))


}