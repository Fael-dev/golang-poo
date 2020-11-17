package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso!"
	}else{
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) ( string, float64 ){
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return "Depósito realizado com sucesso.", c.Saldo
	}else{
		return "Digite um valor maior que zero.", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool{
	if valorTransferencia < c.Saldo && valorTransferencia > 0{
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}else{
		return false
	}
}
