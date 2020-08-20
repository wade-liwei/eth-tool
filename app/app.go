package app

import (
	"errors"
	"sync"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
	"github.com/wade-liwei/eth-tool/contracts"
)

const (
	WaittingNonceDuration = time.Second * 10
)

//var loggerWithApp = logger.With("module", "app")
type App struct {
	sync.Mutex
	EthCliConn       *ethclient.Client
	EthForeignBridge *contracts.ForeignBridge
	EthAuthor        *bind.TransactOpts

	EthRpcClient *rpc.Client
	//IPC
	IpcClient       *rpcclient.Client
	IpcSignerClient *rpcclient.Client
	Config
}

type Config struct {
	BatchSize  uint64
	BlockRange uint64
	//foreign.request_timeout
	ForeignRequestTimeout time.Duration
	//poll_interval
	ForeignPollInterval time.Duration
	PollInterval        time.Duration
	//confirmations
	ForeignConfirmations uint64
	Confirmations        uint64
}

func NewApp(v *viper.Viper, ethAddrStr, ipcAddrStr string, ethContractAddrStr common.Address) (*App, error) {
	var app App

	//ethAddrStr := v.GetString("foreign.eth_rpc_addr")
	//ethContractAddrStr := v.GetString("foreign.foreign_contract_address")
	ethPrvKeyStr := v.GetString("foreign.prv_key")
	//ipcAddrStr := v.GetString("home_bridge.ipc_rpc_addr")
	ipcRpcUserStr := v.GetString("home_bridge.ipc_rpc_user")
	ipcRpcPassStr := v.GetString("home_bridge.ipc_rpc_pass")

	if err := app.NewEthRpcClientAndForeignBridge(ethAddrStr, ethContractAddrStr); err != nil {
		return nil, err
	}

	//write permission
	app.EthAuthor = SetPrvKey(ethPrvKeyStr)
	if err := app.NewIpcClient(ipcAddrStr, ipcRpcUserStr, ipcRpcPassStr); err != nil {
		return nil, err
	}

	ipcSignerAddrStr := v.GetString("home_bridge.ipc_rpc_signer_addr")
	ipcSignerRpcUserStr := v.GetString("home_bridge.ipc_rpc_signer_user")
	ipcSignerRpcPassStr := v.GetString("home_bridge.ipc_rpc_signer_pass")

	if ipcSignerAddrStr != "" && ipcSignerRpcUserStr != "" && ipcSignerRpcPassStr != "" {
		if err := app.NewIpcSignerClient(ipcSignerAddrStr, ipcSignerRpcUserStr, ipcSignerRpcPassStr); err != nil {
			return nil, err
		}
	}
	app.ForeignConfirmations = uint64(v.GetInt64("foreign.required_confirmations"))
	app.Confirmations = app.ForeignConfirmations
	app.ForeignPollInterval = v.GetDuration("poll_interval")
	app.PollInterval = app.ForeignPollInterval
	batchSize, err := app.EthForeignBridge.BatchSize(nil)
	if err != nil {
		return nil, err
	}

	app.BatchSize = batchSize.Uint64()

	blockRange, err := app.EthForeignBridge.BlockRange(nil)
	if err != nil {
		return nil, err
	}

	app.BlockRange = blockRange.Uint64()

	minBalance, err := app.EthForeignBridge.MinBalance(nil)

	_ = minBalance

	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (a *App) NewIpcSignerClient(host, user, pass string) error {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		//Host:         "192.168.1.113:18332",
		Host: host,
		//User:         "rpcuser",
		User: user,
		//Pass:         "rpcpswd",
		Pass:         pass,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return err
	}

	num, err := client.GetInfo()
	if err != nil {
		return err
	}

	_ = num

	a.IpcSignerClient = client
	return nil
}

func (a *App) NewIpcClient(host, user, pass string) error {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		//Host:         "192.168.1.113:18332",
		Host: host,
		//User:         "rpcuser",
		User: user,
		//Pass:         "rpcpswd",
		Pass:         pass,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return err
	}

	num, err := client.GetInfo()
	if err != nil {
		return err
	}

	_ = num

	a.IpcClient = client
	return nil
}

//
//   ETH
//
func (a *App) NewEthRpcClientAndForeignBridge(url string, contractAddr common.Address) error {
	if err := a.NewEthRpcClient(url); err != nil {
		return err
	}
	return a.NewEthForeignBridge(url, contractAddr)
}

//query eth e.g. current block number
func (a *App) NewEthRpcClient(url string) error {
	client, err := rpc.Dial(url)
	if err != nil {
		return err
	}
	a.EthRpcClient = client
	return nil
}

//query contract  e.g. totalSupply
//and write without permission  e.g. transfer
func (a *App) NewEthForeignBridge(url string, contractAddr common.Address) error {
	conn, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	a.EthCliConn = conn
	foreignBridge, err := contracts.NewForeignBridge(contractAddr, conn)
	if err != nil {
		return err
	}
	a.EthForeignBridge = foreignBridge

	return nil
}

/*
EthCliConn       *ethclient.Client
EthForeignBridge *contracts.ForeignBridge
Author           *bind.TransactOpts

EthRpcClient *rpc.Client
*/
func (a *App) GetEthConn() (*ethclient.Client, *contracts.ForeignBridge, *bind.TransactOpts, *rpc.Client, error) {

	if a.EthCliConn == nil {
		return nil, nil, nil, nil, errors.New("EthCliConn is nil, check app init eth")
	}

	if a.EthForeignBridge == nil {
		return nil, nil, nil, nil, errors.New("EthForeignBridge is nil, check app init eth")
		return nil, nil, nil, nil, errors.New("Author is nil, check app init eth")
	}
	if a.EthAuthor == nil {
	}
	if a.EthRpcClient == nil {
		return nil, nil, nil, nil, errors.New("EthRpcConn is nil, check app init eth")
	}
	return a.EthCliConn, a.EthForeignBridge, a.EthAuthor, a.EthRpcClient, nil
}

func (a *App) GetIpcConn() (*rpcclient.Client, error) {

	if a.IpcClient == nil {
		return nil, errors.New("Ipc client is nil,please check app init ipc")
	}
	return a.IpcClient, nil
}
