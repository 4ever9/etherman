/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// nonceCmd represents the nonce command
var nonceCmd = &cobra.Command{
	Use:   "nonce",
	Short: "Get the nonce for the account",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("lack of account address")
		}

		rpcUrl := viper.GetString("rpc-url")
		client := req.C()
		resp, err := client.R().
			SetBody(map[string]any{
				"method":  "eth_getTransactionCount",
				"params":  []string{args[0]},
				"id":      1,
				"jsonrpc": "2.0",
			}).
			Post(rpcUrl)
		if err != nil {
			return err
		}

		fmt.Println(resp)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(nonceCmd)
}
