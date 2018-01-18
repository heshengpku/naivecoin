package transaction

type Transcation struct {
	ID   string `json:"id"`   // random id (64 bytes)
	Data Txs    `json:"data"` // transcations data
	Hash string `json:"hash"` // hash taken from the contents of the transaction: sha256 (id + data) (64 bytes)
	Type string `json:"type"` // transaction type (regular, fee, reward)
}

type Txs struct {
	Inputs  []InPut  `json:"inputs"`  // Transaction inputs
	Outputs []OutPut `json:"outputs"` // Transaction outputs
}

type InPut struct {
	Tx        string `json:"transaction"` // transaction hash taken from a previous unspent transaction output (64 bytes)
	Index     int    `json:"index"`       // index of the transaction taken from a previous unspent transaction output
	Amount    int    `json:"amount"`      // amount of satoshis
	Address   string `json:"address"`     // from address (64 bytes)
	Signature string `json:"signature"`   // transaction input hash: sha256 (transaction + index + amount + address) signed with owner address's secret key (128 bytes)
}

type OutPut struct {
	Amount  int    `json:"amount"`    // amount of satoshis
	Address string `json:"signature"` // to address (64 bytes)
}

// func (tx Transcation) toHash() string {
// 	return utils.CalculateHash(tx.ID + fmt.Sprint(tx.Data))
// }

// func (tx InPut) toHash() string {
// 	return utils.CalculateHash(tx.Tx + tx.Index)
// }

// func (tx Transcation) check() error {
// 	if tx.Hash != tx.toHash {
// 		return errors.New("Invalid transaction hash")
// 	}

// 	for _, input := range tx.Data.Inputs {

// 	}
// }
