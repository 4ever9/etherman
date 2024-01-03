/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/4ever9/etherman/types"
	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

// chainIdCmd represents the chainId command
var chainIdCmd = &cobra.Command{
	Use:   "chain-id",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		rpcUrl := viper.GetString("rpc-url")
		client := req.C()
		resp, err := client.R().
			SetBody(map[string]any{
				"method":  "eth_chainId",
				"params":  []string{},
				"id":      1,
				"jsonrpc": "2.0",
			}).
			Post(rpcUrl)
		if err != nil {
			return err
		}

		isRaw := viper.GetBool("raw")
		if isRaw {
			fmt.Println(resp)
			return nil
		}

		var jr types.JsonrpcResponse
		if err := resp.UnmarshalJson(&jr); err != nil {
			return err
		}

		nonce, _ := new(big.Int).SetString(strings.Replace(jr.Result, "0x", "", -1), 16)

		fmt.Println(nonce.Uint64())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(chainIdCmd)
}
