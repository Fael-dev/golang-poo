package main

import (
	"fmt"
	c "golang-poo/contas"
)

func main(){
	conta1 := c.ContaCorrente{Titular: "Rafael", Saldo: 25.80}
	conta2 := c.ContaCorrente{Titular: "Beatriz", Saldo: 2900.80}

	fmt.Println("=========================== Valores antes da transferencia ===========================")
	fmt.Println(conta1)
	fmt.Println(conta2)
	fmt.Println("=========================== Valores depois da transferencia ===========================")
	status := conta2.Transferir(1000, &conta1)
	fmt.Println(conta1)
	fmt.Println(conta2)
	fmt.Println(status)
}