package fabric

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/beatoz-ryan/bprn-sdk-go/types"
	"github.com/holiman/uint256"
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	"github.com/hyperledger/fabric-protos-go-apiv2/peer"
)

type BaseChaincode struct {
	chaincodeName string
}

// TODO : 다른 체인코드에서 InitLedger를 호출했을때 결과 확인
func (bc *BaseChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	ledger := NewBaseChaincodeLedger(ctx)

	selfChaincodeName, err := bc.CallerChaincodeName(ctx)
	if err != nil {
		return fmt.Errorf("failed to get caller chaincode name: %w", err)
	}
	bc.chaincodeName = selfChaincodeName

	if err := ledger.PutSelfChaincodeName(selfChaincodeName); err != nil {
		return fmt.Errorf("failed to put self chaincode name: %w", err)
	}

	fmt.Println("[BaseChaincode.InitLedger] selfChaincodeName: ", bc.ChaincodeName())

	return nil
}

func (bc *BaseChaincode) ChannelName(ctx contractapi.TransactionContextInterface) string {
	return ctx.GetStub().GetChannelID()
}

func (bc *BaseChaincode) ChaincodeAddress(ctx contractapi.TransactionContextInterface) (string, error) {
	ledger := NewBaseChaincodeLedger(ctx)
	selfChaincodeName, err := ledger.GetSelfChaincodeName()
	if err != nil {
		return "", err
	}

	return GenerateChaincodeAddress(ctx.GetStub().GetChannelID(), selfChaincodeName), nil
}

func (bc *BaseChaincode) ChaincodeNameFromLedger(ctx contractapi.TransactionContextInterface) (string, error) {
	ledger := NewBaseChaincodeLedger(ctx)
	selfChaincodeName, err := ledger.GetSelfChaincodeName()
	if err != nil {
		return "", err
	}
	bc.chaincodeName = selfChaincodeName
	return selfChaincodeName, nil
}

func (bc *BaseChaincode) ChaincodeName() string {
	return bc.chaincodeName
}

func (bc *BaseChaincode) GetChainId(ctx contractapi.TransactionContextInterface) (uint256.Int, error) {
	return GetChainId(ctx.GetStub())
}

func (bc *BaseChaincode) InvokeChaincode(ctx contractapi.TransactionContextInterface, chaincodeName string, methodName string, methodArgs []string) *peer.Response {
	stub := ctx.GetStub()

	args := [][]byte{[]byte(methodName)}
	for _, methodArg := range methodArgs {
		args = append(args, []byte(methodArg))
	}

	return stub.InvokeChaincode(chaincodeName, args, stub.GetChannelID())
}

func (bc *BaseChaincode) CallerChaincodeName(ctx contractapi.TransactionContextInterface) (string, error) {
	callerChaincodeName, err := NewFabricUtil(ctx.GetStub()).CallerChaincodeName()
	if err != nil {
		return "", fmt.Errorf("failed to get caller chaincode name: %w", err)
	}

	return callerChaincodeName, nil
}

func (bc *BaseChaincode) GetSignerAddress(ctx contractapi.TransactionContextInterface, sig string, methodName string, methodParams []string) (*types.Address, error) {
	ledger := NewBaseChaincodeLedger(ctx)
	selfChaincodeName, err := ledger.GetSelfChaincodeName()
	if err != nil {
		return nil, err
	}

	return bc.SignerAddress(sig, selfChaincodeName, methodName, methodParams)
}

func (bc *BaseChaincode) SignerAddress(sig string, chaincodeName string, methodName string, methodParams []string) (*types.Address, error) {
	return SigVerifyAndSignerAddress(sig, chaincodeName, methodName, methodParams)
}

func (bc *BaseChaincode) SignerAddressFromLedger(ledger *BaseLedger, sig string, methodName string, methodParams []string) (*types.Address, error) {
	chaincodeName, err := ledger.GetSelfChaincodeName()
	if err != nil {
		return nil, err
	}
	return bc.SignerAddress(sig, chaincodeName, methodName, methodParams)
}

func (bc *BaseChaincode) SetEvent(ctx contractapi.TransactionContextInterface, event interface{}) error {
	eventType := reflect.TypeOf(event)
	eventName := eventType.Name()

	// 포인터인 경우 실제 타입 이름 가져오기
	if eventType.Kind() == reflect.Ptr {
		eventName = eventType.Elem().Name()
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to obtain JSON encoding: %v", err)
	}

	err = ctx.GetStub().SetEvent(eventName, eventJSON)
	if err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}

	return nil
}
