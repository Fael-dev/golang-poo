package contas

import (
	"golang-poo/clientes"
)

type ContaPoupanca struct{
	Titular clientes.Titular
	Agencia, NumeroDaConta, Operacao      int
	saldo   float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	}else{
		return "Saldo insuficiente!"
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) ( string, float64 ){
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso.", c.saldo
	}else{
		return "Digite um valor maior que zero.", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64{
	return c.saldo
}
