package generator

//
//func TestErc20ArgsGenerator(t *testing.T) {
//	evmTxParam := types.NewDefaultEvmTxParam()
//	toAddress, err := types.NewAddress("0xAbC1234567890DefABC1234567890dEFaBC12345")
//	if err != nil {
//		t.Fatalf("Failed to create address: %v", err)
//	}
//	amount, err := types.NewAmountFromDecimal("10000")
//	if err != nil {
//		t.Fatalf("Failed to create amount: %v", err)
//	}
//	chaincodeParams := []string{toAddress.ToHexString(), amount.ToHex()}
//	signature := ""
//
//	beatozErc20Param := NewBeatozErc20Param("", evmTxParam, chaincodeParams, signature)
//	args := beatozErc20Param.toArray()
//
//	beatozErc20Param2 := NewBeatozErc20ParamFromArray(args)
//
//	isEqual := beatozErc20Param.isEqual(beatozErc20Param2)
//	if !isEqual {
//		t.Fatalf("BeatozErc20Param not equal after conversion")
//	}
//}
//
//func TestErc20ArgsGenerator_FromAccount(t *testing.T) {
//	fromAccount := types.NewAccount()
//	toAddress, err := types.NewAddress("0xAbC1234567890DefABC1234567890dEFaBC12345")
//	if err != nil {
//		t.Fatalf("Failed to create address: %v", err)
//	}
//	amount, err := types.NewAmountFromDecimal("10000")
//	if err != nil {
//		t.Fatalf("Failed to create amount: %v", err)
//	}
//	chaincodeParams := []string{toAddress.ToHexString(), amount.ToHex()}
//
//	evmTxParam := types.NewDefaultEvmTxParam()
//
//	sigMsgGen := NewSigMsgGenerator()
//	sigMsg, err := sigMsgGen.GenerateSigMsg("erc20", "transfer", chaincodeParams, *evmTxParam)
//	if err != nil {
//		t.Fatalf("Failed to generate sigMsg: %v", err)
//	}
//
//	signer := types.NewSigner()
//	signature, err := signer.Sign(fromAccount, sigMsg)
//	if err != nil {
//		t.Fatalf("Failed to sign sigMsg: %v", err)
//	}
//	signatureHex := hex.EncodeToString(signature)
//
//	beatozErc20Param := NewBeatozErc20Param("", evmTxParam, chaincodeParams, signatureHex)
//	chaincodeArgs := beatozErc20Param.toArray()
//
//	//BeatozChaincodeParamGenerator.generate(fromAccount, chaincodeName, chaincodeMethodName, chaincodeParams)
//
//	beatozErc20Param2 := NewBeatozErc20ParamFromArray(chaincodeArgs)
//
//	isEqual := beatozErc20Param.isEqual(beatozErc20Param2)
//	if !isEqual {
//		t.Fatalf("BeatozErc20Param not equal after conversion")
//	}
//
//	// 클라이언트에서 transfer 호출시, 필요한 인자값들 구성
//}
