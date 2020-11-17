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

func (c *ContaCorrente) Depositar(valorDeposito float64) ( string, float64 ){
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso.", c.saldo
	}else{
		return "Digite um valor maior que zero.", c.saldo
	}
}

func main(){
	conta := ContaCorrente{}
	conta.titular = "Rafael Leonan"
	conta.saldo = 2.79

	status, valor := conta.Depositar(900.0)
	fmt.Println("Status: ", status)
	fmt.Println("Saldo: ", valor)
}