package generator

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"linker-chaincodes/types"
	"testing"

	"github.com/beatoz/beatoz-go/types/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func Test2(t *testing.T) {
	hexSig := "a252b52764123787625a7420c41092d15e94b789ac86454caf32c22fe052c4434a2f789ab481e2cec4d8bdc3b7aa63417582807cde986ada94b7c499814d5e9c01"
	sigBytes, err := hex.DecodeString(hexSig)

	if err != nil {
		t.Fatalf("failed to decode sigMsg: %e", err)
	}
	t.Logf("sigMsg: %s", sigBytes)

	sigMsgGenerator := NewSigMsgGenerator()
	sigMsg, err := sigMsgGenerator.GenerateSigMsg("new-erc20-aaa", "InitLedger", []string{"", "KRW Beatoz", "KRWB", "18", "100000000000"})
	if err != nil {
		t.Fatalf("Failed to generate sigMsg: %v", err)
	}

	sigMsgHex := hex.EncodeToString(sigMsg)
	fmt.Println("sigMsgHex:", sigMsgHex)

	//crypto.Sign(sigMsg)

	sigVerifier := NewSigVerifier()
	address, err := sigVerifier.VerifySignature(sigMsg, hexSig)
	if err != nil {
		t.Fatalf("failed to verify signature: %e", err)
	}

	fmt.Println("address: ", address)
}

func TestNewSigVerifier(t *testing.T) {
	ccName := "erc20"
	//methodName := "transfer"
	methodName := "InitLedger"
	evmTxParam := types.NewEvmTxParam(1, 0, 21000, *uint256.NewInt(100000000))
	evmTxParam.ToArray()

	sigMsgGenerator := NewSigMsgGenerator()
	sigMsg, err := sigMsgGenerator.GenerateSigMsg(
		ccName,
		methodName,
		[]string{"cb4360dc62dd2ff32a0da8643c7cc0def7f2e075", "1000000000000000000"},
	)
	if err != nil {
		t.Fatalf("Failed to generate sigMsg: %v", err)
	}

	privateKeyHex := "5d28588923deee55af4aaa3e54ddef8158ab63323e769b767c4ecc026b595c3f"
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		t.Fatalf("failed to parse private key: %e", err)
	}
	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex)
	if err != nil {
		t.Fatalf("failed to parse private key: %e", err)
	}

	compressedPubKey := crypto.CompressPubkey(privateKey.Public().(*ecdsa.PublicKey))
	hexCompressPubKey := hex.EncodeToString(compressedPubKey)
	fmt.Println(hexCompressPubKey)

	prvKey1, err := crypto.ImportPrvKey(privateKeyBytes)
	if err != nil {
		t.Fatalf("failed to import private key: %e", err)
	}
	prvKey2, err := crypto.ImportPrvKeyHex(privateKeyHex)
	if err != nil {
		t.Fatalf("failed to import private key: %e", err)
	}

	address1 := crypto.Pub2Addr(prvKey1.Public().(*ecdsa.PublicKey))
	address2 := crypto.Pub2Addr(prvKey2.Public().(*ecdsa.PublicKey))
	fmt.Println(address1.String())
	fmt.Println(address2.String())

	pubKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.Pub2Addr(pubKey)
	t.Logf("address: %s", address.String())

	// =============================
	//sigMsgHex := hex.EncodeToString(sigMsg)
	sigMsgHex := "f8b0b882376434323932623635313439393332633838663938306439326535306161643866323631613738353566336130316462343230383839393034313833336538623333343966303164646631643935383632646430656662613764353331323230303363363232636338346666336531623563303662666466313365643234626330308a696e69744c6564676572808a4b525720426561746f7a844b5257428231388c313030303030303030303030"
	sigMsg2, err := hex.DecodeString(sigMsgHex)
	if err != nil {
		t.Fatalf("failed to decode sigMsg: %e", err)
	}
	t.Logf("sigMsg: %s", sigMsgHex)

	sig, err := crypto.Sign(sigMsg2, privateKey)
	//crypto.VerifySig()
	sigToAddr, compressedPubKey, err := crypto.Sig2Addr(sigMsg2, sig)
	if err != nil {
		t.Fatalf("failed to verify signature: %e", err)
	}

	_ = sigToAddr
	_ = compressedPubKey
	_ = sigMsg
}
