package cmd

import (
	// "fmt"
	// "os"
	//"github.com/spf13/viper"
	// homedir "github.com/mitchellh/go-homedir"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.PersistentFlags().StringVar(&prvkeyHex, "prvkeyHex", "", "please provide your eth private key.")
}

var prvkeyHex string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "eth-tool",
	Short: "test web3",
	Long:  `batch eth txs test`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//
//
//
// import (
// 	"fmt"
// 	"os"
// 	homedir "github.com/mitchellh/go-homedir"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
//
// )
//
// var logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
// var loggerWithCmd log.Logger
// var cfgFile string
// var prvkeyHex string
//
// // RootCmd represents the base command when called without any subcommands
// var RootCmd = &cobra.Command{
// 	Use:   "bridge",
// 	Short: "Bridge between IPC and  ethereum-based networks",
// 	Long:  `ipc-bridge is currently an ERC20 token contract on one ethereum-based blockchain`,
//
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//config
// 		v := viper.GetViper()
// 		log_level_str := v.GetString("log_level")
// 		//loggerWithCmd.Debug("read log level", "logLevel", log_level_str)
//
// 		logger, _ = tmflags.ParseLogLevel(log_level_str, logger, "*")
// 		loggerWithCmd = logger.With("module", "cmd")
//
// 		//dbPath := v.GetString("db_path")
// 		dbPath := GetDbPath()
// 		loggerWithCmd.Debug("read bridge idx from db", "dbPath", dbPath)
// 		db, err := app.GetIpcEthDbInstance(dbPath)
// 		if err != nil {
// 			logger.Error("get ipc eth db instance ", "err", err.Error())
// 			return
// 		}
//
// 		foreignContract, homeMultiAddr, err := db.GetForeignContractAddressAndHomeMultiAddr()
// 		if err != nil {
// 			loggerWithCmd.Error("ipc eth db instance GetForeignContractAddress", "err", err.Error())
// 			return
// 		}
//
// 		ethAddrStr := v.GetString("foreign.eth_rpc_addr")
// 		ipcAddrStr := v.GetString("home_bridge.ipc_rpc_addr")
//
// 		loggerWithCmd.Info("bridge Establishing ipc and eth connection", "ethRpcAddr",
// 			ethAddrStr, "foreignContract", foreignContract.Hex(),
// 			"ipcRpcAddr", ipcAddrStr, "ipcMultiSignAddr", homeMultiAddr)
//
// 		//TODO  用户名密码需要参考tendermint或想个办法实现
// 		appInstance, err := app.NewApp(v, ethAddrStr, ipcAddrStr, foreignContract, logger)
//
// 		if err != nil {
// 			loggerWithCmd.Error("app init", "err", err.Error())
// 			return
// 		}
//
// 		b, err := bridge.NewBridge(appInstance, db, logger)
// 		if err != nil {
// 			loggerWithCmd.Error("new bridge err", "error", err.Error())
// 			return
// 		}
//
// 		if err := b.Stream(); err != nil {
// 			loggerWithCmd.Error("bridge stream run", "error", err.Error())
// 			return
// 		}
// 	},
// }
//
// // Execute adds all child commands to the root command and sets flags appropriately.
// // This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	if err := RootCmd.Execute(); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }
//
// func init() {
// 	cobra.OnInitialize(initConfig, initDbConfig, initDeployConfig)
// 	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bridge.yaml)")
// 	RootCmd.PersistentFlags().StringVar(&prvkeyHex, "prvkeyHex", "", "please provide your eth private key.")
// }
//
// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
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
// 		viper.SetConfigName(".ipcBridge")
// 	}
//
// 	viper.AutomaticEnv() // read in environment variables that match
//
// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
