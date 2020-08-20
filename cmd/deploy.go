package cmd

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/wade-liwei/eth-tool/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wade-liwei/eth-tool/app"
)

var deployCfgFile string

func init() {
	RootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().StringVar(&deployCfgFile, "dpConfig", "", "deploy config file (default is $HOME/.ipcBridgeIdx.toml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy  contract to evm chain.",
	Run: func(cmd *cobra.Command, args []string) {

		if prvkeyHex == "" {
			fmt.Println("please provide your private key.")
			return
		}

		ethRpcAddrStr := viper.GetString("eth_rpc_addr")

		deploy := app.NewDeploy(prvkeyHex)
		if err := deploy.Dial(ethRpcAddrStr); err != nil {
			fmt.Println(err)
			return
		}

		tokenNameStr := viper.GetString("token_name")
		symbolStr := viper.GetString("symbol")
		decimalsUint8 := uint8(viper.GetInt64("decimal"))
		batchSize := big.NewInt(viper.GetInt64("batch_size"))
		blockRange := big.NewInt(viper.GetInt64("block_range"))
		minBalance := big.NewInt(viper.GetInt64("min_balance"))

		authors := []common.Address{}

		authorsAsStr := viper.GetStringSlice("authorities")

		for _, v := range authorsAsStr {
			authors = append(authors, common.HexToAddress(v))
		}

		req := big.NewInt(viper.GetInt64("requiredSignatures"))

		addr, err := deploy.Deploy(contracts.ForeignBridgeABI, contracts.ForeignBridgeBin, tokenNameStr, symbolStr, decimalsUint8, authors, req, batchSize, blockRange, minBalance)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("deploy addr:  "+addr.String(), "tokenName", tokenNameStr, "symbol", symbolStr, "decimal", decimalsUint8, "required", req, "batchSize", batchSize, "blockRange", blockRange, "minBalance", minBalance)

		time.Sleep(time.Second * 5)

		appInstance := app.App{}

		err = appInstance.NewEthRpcClientAndForeignBridge(ethRpcAddrStr, *addr)

		if err != nil {
			return
		}

		appInstance.EthAuthor = app.SetPrvKey(prvkeyHex)
		//opts *bind.TransactOpts, _requiredSignatures *big.Int, _authorities []common.Address, _name string, _symbol string, _decimals uint8
		// txHash, err := appInstance.EthForeignBridge.SetForeignBridge(appInstance.EthAuthor, req, authors, tokenNameStr, symbolStr, decimalsbig)
		//
		// if err != nil {
		// 	logger.Fatal(err)
		// }
		//
		// logger.Debugf("set foregin bridge  tx hash: %v ", txHash.Hash().String())
		//
		// logger.Debugf("authors: %v, req: %v  tokenName: %v  symbolStr: %v  decimal: %v", authorsAsStr, req.String(), tokenNameStr, symbolStr, decimalsUint8)
		//
		// ctx := context.Background()
		//
		// receipt, err := bind.WaitMined(ctx, appInstance.EthCliConn, txHash)
		// if err != nil {
		// 	logger.Fatalf("deploy foreign bridge err: %v", err)
		// }
		//
		// if receipt.Status != 1 {
		// 	logger.Fatal("receipt.Status != 1")
		// 	//return errors.New(fmt.Sprintf("submit msg to eth contract, expect receipt status ==1 but actual:%d", receipt.Status))
		// }

		client, err := rpc.Dial(ethRpcAddrStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		blockNum, err := app.GetEthBlockNumByRpc(client)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("foreign chain deployed addr:%s  blockNum: %d \n ", addr.String(), blockNum)
	},
}

//
// import (
// 	"fmt"
// 	"math/big"
// 	"os"
// 	"time"
//
// 	"192.168.1.188/ipc/ipc-bridge/app"
// 	"192.168.1.188/ipc/ipc-bridge/contracts"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/rpc"
// 	homedir "github.com/mitchellh/go-homedir"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )
//
// var deployCfgFile string
//
// func init() {
// 	RootCmd.AddCommand(deployCmd)
// 	deployCmd.PersistentFlags().StringVar(&deployCfgFile, "dpConfig", "", "deploy config file (default is $HOME/.ipcBridgeIdx.toml)")
// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
//
// var deployCmd = &cobra.Command{
// 	Use:   "deploy",
// 	Short: "deploy  contract to geth poa chain.",
// 	Run: func(cmd *cobra.Command, args []string) {
//
// 		if prvkeyHex == "" {
// 			fmt.Println("please provide your private key.")
// 			return
// 		}
//
// 		ethRpcAddrStr := viper.GetString("eth_rpc_addr")
//
// 		deploy := app.NewDeploy(prvkeyHex)
// 		if err := deploy.Dial(ethRpcAddrStr); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
//
// 		tokenNameStr := viper.GetString("token_name")
// 		symbolStr := viper.GetString("symbol")
// 		decimalsUint8 := uint8(viper.GetInt64("decimal"))
// 		batchSize := big.NewInt(viper.GetInt64("batch_size"))
// 		blockRange := big.NewInt(viper.GetInt64("block_range"))
// 		minBalance := big.NewInt(viper.GetInt64("min_balance"))
//
// 		authors := []common.Address{}
//
// 		authorsAsStr := viper.GetStringSlice("authorities")
//
// 		for _, v := range authorsAsStr {
// 			authors = append(authors, common.HexToAddress(v))
// 		}
//
// 		req := big.NewInt(viper.GetInt64("requiredSignatures"))
//
// 		addr, err := deploy.Deploy(contracts.ForeignBridgeABI, contracts.ForeignBridgeBin, tokenNameStr, symbolStr, decimalsUint8, authors, req, batchSize, blockRange, minBalance)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
//
// 		logger.Debug("deploy addr:  "+addr.String(), "tokenName", tokenNameStr, "symbol", symbolStr, "decimal", decimalsUint8, "required", req, "batchSize", batchSize, "blockRange", blockRange, "minBalance", minBalance)
//
// 		time.Sleep(time.Second * 5)
//
// 		appInstance := app.App{}
//
// 		err = appInstance.NewEthRpcClientAndForeignBridge(ethRpcAddrStr, *addr)
//
// 		if err != nil {
// 			logger.Error("init eth rpc and contract", "error", err.Error())
// 			return
// 		}
//
// 		appInstance.EthAuthor = app.SetPrvKey(prvkeyHex)
// 		//opts *bind.TransactOpts, _requiredSignatures *big.Int, _authorities []common.Address, _name string, _symbol string, _decimals uint8
// 		// txHash, err := appInstance.EthForeignBridge.SetForeignBridge(appInstance.EthAuthor, req, authors, tokenNameStr, symbolStr, decimalsbig)
// 		//
// 		// if err != nil {
// 		// 	logger.Fatal(err)
// 		// }
// 		//
// 		// logger.Debugf("set foregin bridge  tx hash: %v ", txHash.Hash().String())
// 		//
// 		// logger.Debugf("authors: %v, req: %v  tokenName: %v  symbolStr: %v  decimal: %v", authorsAsStr, req.String(), tokenNameStr, symbolStr, decimalsUint8)
// 		//
// 		// ctx := context.Background()
// 		//
// 		// receipt, err := bind.WaitMined(ctx, appInstance.EthCliConn, txHash)
// 		// if err != nil {
// 		// 	logger.Fatalf("deploy foreign bridge err: %v", err)
// 		// }
// 		//
// 		// if receipt.Status != 1 {
// 		// 	logger.Fatal("receipt.Status != 1")
// 		// 	//return errors.New(fmt.Sprintf("submit msg to eth contract, expect receipt status ==1 but actual:%d", receipt.Status))
// 		// }
//
// 		client, err := rpc.Dial(ethRpcAddrStr)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
//
// 		blockNum, err := app.GetEthBlockNumByRpc(client)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
//
// 		fmt.Printf("foreign chain deployed addr:%s  blockNum: %d \n ", addr.String(), blockNum)
// 	},
// }
//
// // initConfig reads in config file and ENV variables if set.
// func initDeployConfig() {
//
// 	if deployCfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(deployCfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
//
// 		// Search config in home directory with name ".bridge" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".deployConfig")
// 	}
//
// 	viper.AutomaticEnv() // read in environment variables that match
//
// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
