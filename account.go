package query

import (
	// "errors"
	// "math/big"
	// "strconv"
    // "fmt"

	"encoding/json"
//	"fmt"

	//"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/thetatoken/theta/utils"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/rpc"

	rpcc "github.com/ybbus/jsonrpc"
)

func Account(address string) string {
	
	client := rpcc.NewRPCClient(viper.GetString("remoteRPCEndpoint"))

	res, err := client.Call("theta.GetAccount", rpc.GetAccountArgs{
		Address: address,
		Height:  common.JSONUint64(uint64(0)),
		Preview: false})
	if err != nil {
		utils.Error("Failed to get account details: %v\n", err)
	}
	if res.Error != nil {
		utils.Error("Failed to get account details: %v\n", res.Error)
	}
	json, err := json.MarshalIndent(res.Result, "", "    ")
	if err != nil {
		utils.Error("Failed to parse server response: %v\n%v\n", err, string(json))
	}
	// fmt.Println(string(json))
	// fmt.Println("ep:" + endpoint + " addr:" + address)
	return(string(json))
}


