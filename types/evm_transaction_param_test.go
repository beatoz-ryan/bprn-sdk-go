package types

import (
	"testing"
)

func TestEvmTransactionParam_Serialize(t *testing.T) {
	evmTxParam := NewDefaultEvmTxParam()
	serialized, err := evmTxParam.ToArray()
	if err != nil {
		t.Errorf("EvmTransactionParam ToArray() failed")
	}
	_ = serialized

	//evmTxParam2 := NewDefaultEvmTxParam()
	//evmTxParam2.From(arr)
	//result := evmTxParam.IsEqual(*evmTxParam2)
	//if !result {
	//	t.Errorf("EvmTransactionParam ToArray() and From() failed")
	//}
}
