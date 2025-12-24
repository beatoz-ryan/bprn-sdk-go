package generator

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/beatoz/beatoz-go/types/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func TestSigMsg(t *testing.T) {
	chaincodeName := "new-erc20-3"
	methodName := "InitLedger"
	chaincodeParams := []string{"", "KRW Beatoz", "KRWB", "18", "100000000000"}
	sigMsgGenerator := NewSigMsgGenerator()
	sigMsg, err := sigMsgGenerator.GenerateSigMsg(
		chaincodeName,
		methodName,
		chaincodeParams,
	)
	if err != nil {
		t.Fatalf("Failed to generate sigMsg: %v", err)
	}

	hexSigMsg := hex.EncodeToString(sigMsg)
	fmt.Println(hexSigMsg)

	srcSigMsgHex := "f8388b6e65772d65726332302d338a496e69744c6564676572808a4b525720426561746f7a844b5257428231388c313030303030303030303030"
	if hexSigMsg != srcSigMsgHex {
		t.Fatalf("sigMsg is not correct: %s", hexSigMsg)
	}

	//const fromSigner = new Account("cb4360dc62dd2ff32a0da8643c7cc0def7f2e075", "5d28588923deee55af4aaa3e54ddef8158ab63323e769b767c4ecc026b595c3f")

	srcAddress := "cb4360dc62dd2ff32a0da8643c7cc0def7f2e075"
	privateKeyHex := "5d28588923deee55af4aaa3e54ddef8158ab63323e769b767c4ecc026b595c3f"
	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex)
	if err != nil {
		t.Fatalf("failed to parse private key: %e", err)
	}
	sig, err := crypto.Sign(sigMsg, privateKey)

	srcHexSig := "ad3fca2a1be115059b6c8f0318e2589ecce2ad81d7f5ee9ce22efe3062f67931740892051e94943ec6b6e813f40ba5f130db010122dd804c92cf8789c58f159801"
	hexSig := hex.EncodeToString(sig)
	if hexSig != srcHexSig {
		t.Fatalf("sig is not correct: %s", hexSig)
	}
	fmt.Println("hexSig: ", hexSig)

	sigVerifier := NewSigVerifier()
	address, err := sigVerifier.VerifySignature2(sigMsg, sig)
	if err != nil {
		t.Fatalf("failed to verify signature: %e", err)
	}

	if srcAddress != address.String() {
		t.Fatalf("address mismatch: expected %s, got %s", srcAddress, address.String())
	}
	fmt.Printf("Signature verification successful: address = %s\n", address.String())
}

//
//import (
//	"encoding/hex"
//	"linker-chaincodes/fabric"
//	"linker-chaincodes/types"
//	"testing"
//
//	"github.com/beatoz/beatoz-go/types/crypto"
//	"github.com/holiman/uint256"
//)
//
//func TestNewSigMsgGenerator(t *testing.T) {
//	sigMsgGenerator := NewSigMsgGenerator()
//
//	ccName := "linker-endpoint-8"
//	methodName := "InitLedger"
//	sig := "97b0496aace9a4da0ffcd59856dd46e6d70cdabf23c71254376708e3fb4a3cde1008625483fca2f505208a76048eba8691d33fee4f60317fc5dca47c0e97f53300"
//	address, err := fabric.SigVerifyAndSignerAddress(sig, ccName, methodName, []string{})
//	if err != nil {
//		t.Fatalf("Failed to verify sig: %v", err)
//	}
//	_ = address
//
//	s1 := uint256.NewInt(100000000)
//	hex1 := s1.Hex()
//	hex2 := hex.EncodeToString(s1.Bytes())
//	_, _ = hex1, hex2
//	s2, _ := hex.DecodeString(hex1)
//	s3, _ := hex.DecodeString("5f5e100")
//	s4, _ := hex.DecodeString("05f5e100")
//	//s3 := []byte("05f5e100")
//	s5 := []byte(hex1)
//	_ = s1
//	_ = s2
//	_ = s3
//	_ = s4
//	_ = s5
//
//	evmTxParam := types.NewEvmTxParam(1, 0, 21000, *uint256.NewInt(100000000))
//	evmTxParam.ToArray()
//
//	sigMsg, err := sigMsgGenerator.GenerateSigMsg(
//		evmTxParam,
//		"erc20",
//		"transfer",
//		[]string{"cb4360dc62dd2ff32a0da8643c7cc0def7f2e075", "1000000000000000000"},
//	)
//	if err != nil {
//		t.Fatalf("Failed to generate sigMsg: %v", err)
//	}
//
//	sigMsgHex := hex.EncodeToString(sigMsg)
//	t.Logf("sigMsg: %s", sigMsgHex)
//
//	btzAddress, _, err := crypto.Sign(sigMsg)
//
//}
//
////func TestSigMsgGenerator_GenerateSigMsg(t *testing.T) {
////	types args struct {
////		nonce    string
////		playerId string
////		gameId   string
////	}
////	tests := []struct {
////		name string
////		args args
////		want string
////	}{
////		// TODO: Add test cases.
////	}
////	for _, tt := range tests {
////		t.Run(tt.name, func(t *testing.T) {
////			s := &SigMsgGenerator{}
////			if got := s.GenerateSigMsg(tt.args.nonce, tt.args.playerId, tt.args.gameId); got != tt.want {
////				t.Errorf("GenerateSigMsg() = %v, want %v", got, tt.want)
////			}
////		})
////	}
////}
////
////func TestSigMsgGenerator_GenerateSigMsg2(t *testing.T) {
////	s := &SigMsgGenerator{}
////	sigmsg := s.GenerateSigMsg("from1", "to1", "100000")
////}
