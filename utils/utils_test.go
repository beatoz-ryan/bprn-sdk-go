package utils

import (
	"testing"
)

func TestUtils(t *testing.T) {
	str := "mychannel0"
	//str := "516977165748044443249712"
	chainId, err := StringToUint256(str)
	if err != nil {
		t.Error(err)
	}
	t.Log(chainId)

	chainIdHex := chainId.Hex()
	t.Log(chainIdHex)
	// 726561642074686973
	// 0x6d796368616e6e656c30

	chainIdStr := Uint256ToString(chainId)
	if err != nil {
		t.Error(err)
	}
	t.Log(chainIdStr)
}
