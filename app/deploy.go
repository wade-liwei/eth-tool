package app

import (
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"
)

// Deploy deploys smart contract with abi and bin
type Deploy struct {
	con  *ethclient.Client
	auth *bind.TransactOpts
}

// NewDeploy creates Deployer
func NewDeploy(prvkeyHex string) *Deploy {
	return &Deploy{
		auth: SetPrvKey(prvkeyHex),
	}
}

// Dial creates ethclient with specified URL
func (d *Deploy) Dial(rawURL string) error {
	con, err := ethclient.Dial(rawURL)
	if err != nil {
		return err
	}
	d.con = con
	return nil
}

// Deploy deploys smart contract with given abi and bin
func (d *Deploy) Deploy(contractAbi, contractBin, name, symbol string, decimals uint8, authors []common.Address, req, batchSize, blockRange, minBalance *big.Int) (*common.Address, error) {
	parsed, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}
	dataInputs := make([]interface{}, 0, len(parsed.Constructor.Inputs))

	if cap(dataInputs) == 8 {
		dataInputs = append(dataInputs, req)
		dataInputs = append(dataInputs, authors)
		dataInputs = append(dataInputs, name)
		dataInputs = append(dataInputs, symbol)
		dataInputs = append(dataInputs, decimals)
		dataInputs = append(dataInputs, batchSize)
		dataInputs = append(dataInputs, blockRange)
		dataInputs = append(dataInputs, big.NewInt(0).Mul(minBalance, big.NewInt(1000000000000000000)))
	} else {
		for _, i := range parsed.Constructor.Inputs {
			t := reflect.Zero(i.Type.TupleType)
			v := t.Interface()
			if t.Kind() == reflect.Ptr && t.IsNil() {
				elem := reflect.TypeOf(v).Elem()
				v2 := reflect.New(elem)
				v = v2.Interface()
			}
			dataInputs = append(dataInputs, v)
		}
	}

	d.auth.GasPrice = big.NewInt(5)
	d.auth.GasLimit = 210000

	address, _, _, err := bind.DeployContract(d.auth, parsed, common.FromHex(contractBin), d.con, dataInputs...)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
