package contas

import c "golang-poo/clientes"

type ContaCorrente struct {
	Titular             		c.Titular
	NumeroAgencia, NumeroConta  int
	saldo         				float64
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

func (c *ContaCorrente) ObterSaldo() float64{
	return c.saldo
}
