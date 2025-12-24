package fabric

import (
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

const (
	selfChaincodeNameKey = "selfChaincodeName"
)

type BaseLedger struct {
	Ledger *KeyValueLedger
}

func NewBaseChaincodeLedger(ctx contractapi.TransactionContextInterface) *BaseLedger {
	kvLedger := NewKeyValueLedger(ctx)
	return &BaseLedger{
		Ledger: kvLedger,
	}
}

func (bc *BaseLedger) GetBaseChaincodeLedger() *BaseLedger {
	return bc
}

func (bc *BaseLedger) PutSelfChaincodeName(selfChaincodeName string) error {
	err := bc.Ledger.PutString(selfChaincodeNameKey, selfChaincodeName)
	if err != nil {
		return err
	}
	return nil
}

func (bc *BaseLedger) GetSelfChaincodeName() (string, error) {
	selfChaincodeNameBytes, err := bc.Ledger.Get(selfChaincodeNameKey)
	if err != nil {
		return "", err
	}
	return string(selfChaincodeNameBytes), nil
}
