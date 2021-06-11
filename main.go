package main

import (
    "fmt"
)

func main() {
  //mapas
  m := make(map[string]int)
  m["jose"] =14
  m["Pepito"] =20
  m["Pepito1"] =21
  m["Pepito2"] =22


  for i,v :=range m {
    fmt.Println(i,v)
  }

  value, existe := m["jose"]
  fmt.Println(value,existe)
  value, existe = m["josep"]
  fmt.Println(value,existe)


}