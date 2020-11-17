package main

import (
	"fmt"
	c "golang-poo/contas"
	cl "golang-poo/clientes"
)

func main(){
	cliente := cl.Titular{Nome: "Rafael", Cpf: "083.324.094.90", Profissao: "Desenvolvedor"}
	conta := c.ContaCorrente{Titular: cliente, NumeroAgencia: 1234, NumeroConta: 5678, Saldo: 200}
	fmt.Println(conta)
}