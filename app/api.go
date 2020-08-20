package app

import (
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
)

func GetEthBlockNumByRpc(r *rpc.Client) (uint64, error) {

	var blockNumHex string
	err := r.Call(&blockNumHex, "eth_blockNumber")
	if err != nil {
		return 0, err
	}

	blockNumInt64, err := strconv.ParseInt(blockNumHex[2:], 16, 64)
	if err != nil {
		return 0, err
	}

	return uint64(blockNumInt64), nil
}
