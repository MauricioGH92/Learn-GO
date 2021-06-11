package main

import (
    "fmt"
    "strings"
    "reflect"
)

func reverseString(input []string) []string {
    if len(input) == 0 {
        return input
    }
    return append(reverseString(input[1:]), input[0]) 
}

func isPalindromo(word string) bool{
  splitWord := strings.Split(strings.ToUpper(word), "")
  return reflect.DeepEqual(splitWord, reverseString(splitWord))
}

func main() {
  word :="Ama"
  if(isPalindromo(word)){
    fmt.Println("Es palindromo")
  }else{
    fmt.Println("No es palindromo")
  }
}