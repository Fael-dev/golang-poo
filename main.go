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

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool{
	if valorTransferencia < c.saldo && valorTransferencia > 0{
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}else{
		return false
	}
}

func main(){
	conta1 := ContaCorrente{titular: "Rafael", saldo: 25.80}
	conta2 := ContaCorrente{titular: "Beatriz", saldo: 2900.80}

	fmt.Println("=========================== Valores antes da transferencia ===========================")
	fmt.Println(conta1)
	fmt.Println(conta2)
	fmt.Println("=========================== Valores depois da transferencia ===========================")
	status := conta2.Transferir(1000, &conta1)
	fmt.Println(conta1)
	fmt.Println(conta2)
	fmt.Println(status)
}