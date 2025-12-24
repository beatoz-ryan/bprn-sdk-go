package fabric

import (
	"testing"
)

func TestNewSigMsgGenerator(t *testing.T) {
	ccName := "linker-endpoint-8"
	methodName := "InitLedger"
	sig := "60464fe9ba9ef28ee12d0241ca64dd493d19e9337bd87b99794ac659062d398a57a9c9948c9194aadae2acd097b63c592aeb75570b0d58000c03394b5385555301"
	address, err := SigVerifyAndSignerAddress(sig, ccName, methodName, []string{})

	// cb4360dc62dd2ff32a0da8643c7cc0def7f2e075
	hexAddr := address.String()
	if err != nil {
		t.Fatalf("Failed to verify sig: %v", err)
	}
	_, _ = address, hexAddr
}
