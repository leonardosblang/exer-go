package main

import (
	"fmt"
)

func main() {
  for{
  var c int  
  var a int
  var b int
  fmt.Println("O que deseja fazer? 1- somar 2- sair")
  fmt.Scan(&c)
  if c == 1{
  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Println("A soma é: ", a+b)
  } else if c ==2{
    fmt.Println("SAINDO")
    break
  } else{
      fmt.Println("Entre com 1 ou 2")
    }


      }
}
