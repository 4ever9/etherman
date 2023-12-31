/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("lack of account address")
		}

		rpcUrl := viper.GetString("rpc-url")
		client := req.C()
		resp, err := client.R().
			SetBody(map[string]any{
				"method":  "eth_getBalance",
				"params":  []string{args[0], "latest"},
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

		balance, _ := new(big.Int).SetString(strings.Replace(jr.Result, "0x", "", -1), 16)

		fmt.Println(balance.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// balanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
