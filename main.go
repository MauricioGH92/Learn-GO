package main

	
import (
    "fmt"
)

func main() {
  //declaracion de variables
  helloMessage := "Hello"
  worldMessage := "world"


  // Println
  fmt.Println(helloMessage, worldMessage)
  fmt.Println(helloMessage, worldMessage)

  //printf
  nombre:= "Platzi"
  cursos:= 500

  fmt.Printf("%s tiene más de  %d cursos\n",nombre, cursos)
  fmt.Printf("%v tiene más de  %v cursos\n",nombre, cursos)

  //sprintf
  message := fmt.Sprintf("%s tiene más de  %d cursos",nombre, cursos)
  fmt.Println(message)

  //tipo de dato
  fmt.Printf("nombre: %T\n", nombre)
  fmt.Printf("cursos: %T\n", cursos)


}