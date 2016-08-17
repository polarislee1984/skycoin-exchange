package skycoin_interface

type Gateway struct{}

func (gw *Gateway) GetTx(txid string) (Transaction, error) {
	return Transaction{}, nil
}

func (gw *Gateway) GetRawTx(txid string) (string, error) {
	return "skycoin hello world", nil
}

func (gw *Gateway) DecodeRawTx(rawtx string) (Transaction, error) {
	return Transaction{}, nil
}

func (gw *Gateway) InjectTx(tx *Transaction) (string, error) {
	return "new skycoin transaction", nil
}
