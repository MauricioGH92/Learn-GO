package main

import "fmt"

func main() {
  //declaracion de constantes
  const pi float64 = 3.14
  const pi2 = 3.1314
  fmt.Println("pi:",pi)
  fmt.Println("pi2:",pi2)

  //declaracion de variables
  base :=12
  var altura int = 14
  var area int

  fmt.Println(base, altura, area)

  //zero values
  var a int
  var b float64
  var c string
  var d bool
  fmt.Println(a, b, c, d)


  //Area cuadrado
  const baseCuadrado = 10
  areaCuadrado := baseCuadrado * baseCuadrado

  fmt.Println("Area cuadrado: " , areaCuadrado)


  //operadores aritmeticos
  x:=10
  y:=51

  //suma
  result := x+y
  fmt.Println("Suma: ", result)

  //resta
  result = y-x
  fmt.Println("Resta: ", result)

  //multiplicacion
  result = x*y
  fmt.Println("Multiplicacion: ", result)

  //division
  result = y/x
  fmt.Println("Division: ", result)

  //modulo
  result = y%x
  fmt.Println("Modulo: ", result)

  //Incremental
  x++
  fmt.Println("Incremental: ", x)

  //decremental
  x--
  fmt.Println("Decremental: ", x)

}