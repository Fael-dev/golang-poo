package main

import (
	"fmt"
	"golang-poo/clientes"
	c "golang-poo/contas"
)

type verificarConta interface{
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorBoleto float64){
	conta.Sacar(valorBoleto)
}

func main(){

	cliente := clientes.Titular{ Nome: "Rafael", Cpf: "098.889.988-00", Profissao: "Desenvolvedor" }

	contaP := c.ContaPoupanca{ Titular: cliente, NumeroDaConta: 1234, Agencia: 234, Operacao: 90 }
	contaC := c.ContaCorrente{ Titular: cliente, NumeroConta: 3455, NumeroAgencia: 1234}
	contaP.Depositar(200)
	contaC.Depositar(350)

	PagarBoleto(&contaC, 60)

	fmt.Println(contaC.ObterSaldo())

}