package main

	
import (
    "fmt"
)


func main() {
  //Array inmutables
  var array [4] int
  array[0] = 1
  array[1] = 2
  fmt.Println(array,len(array),cap(array))

  //slice mutables
  var slice []int
  slice = append(slice, 2)
  slice = append(slice, 2)
  slice = append(slice, 2)
  fmt.Println(slice,len(slice),cap(slice))

  var slice2 = []int{1,2,3}
  fmt.Println(slice2,len(slice2),cap(slice2))

  slice2 =append(slice2,slice... )

  fmt.Println(array[1:])
  fmt.Println(slice2)



}