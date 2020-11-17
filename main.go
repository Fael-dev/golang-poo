package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	}else{
		return "Saldo insuficiente!"
	}
}

func main(){
	conta := ContaCorrente{}
	conta.titular = "Rafael Leonan"
	conta.saldo = 2.79
	valorSaque := -1.7
	fmt.Println(conta.saldo)
	fmt.Println(conta.Sacar(valorSaque))
	fmt.Println(conta.saldo)
}