package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func SetKeyStore(keyFilePath, passwd string) *bind.TransactOpts {

	keyJsonBytes, err := ioutil.ReadFile(keyFilePath)

	if err != nil {
		fmt.Println(fmt.Errorf("Failed to read the key json file:%s, err: %v ", keyFilePath, err))
		return nil
	}

	auth, err := bind.NewTransactor(strings.NewReader(string(keyJsonBytes)), passwd)
	if err != nil {
		fmt.Println(fmt.Printf("newtransactor err: %v", err))
		return nil
	}

	return auth
}

// SetPrvKey creates keyed-transactor with specified private key.
func SetPrvKey(prvkeyHex string) *bind.TransactOpts {
	keyBytes := common.FromHex(prvkeyHex)
	key := crypto.ToECDSAUnsafe(keyBytes)
	return bind.NewKeyedTransactor(key)
}

// SprintTransaction print Transaction fields
func SprintTransaction(tx *types.Transaction) (string, error) {
	txbody, err := tx.MarshalJSON()
	if err != nil {
		return "", err
	}
	var txjson bytes.Buffer
	if err := json.Indent(&txjson, txbody, "", "\t"); err != nil {
		return "", err
	}

	return string(txjson.Bytes()), nil
}

func SprintReceipt(receipt *types.Receipt) (string, error) {
	txbody, err := receipt.MarshalJSON()
	if err != nil {
		return "", err
	}
	var txjson bytes.Buffer

	if err := json.Indent(&txjson, txbody, "", "\t"); err != nil {
		return "", err
	}
	return string(txjson.Bytes()), nil
}

func OldSprintReceipt(tx *types.Transaction, bck bind.DeployBackend) (string, error) {

	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, bck, tx)
	if err != nil {
		log.Fatalf("tx mining error:%v\n", err)
	}

	receiptAsJson, err := receipt.MarshalJSON()
	if err != nil {
		log.Fatalf("receipt err: %v", err)
	}

	return string(receiptAsJson), nil
}
