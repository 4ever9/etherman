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

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("lack of account address")
		}

		rpcUrl := viper.GetString("rpc-url")
		fmt.Println(rpcUrl)
		req.DevMode()
		client := req.C()
		resp, err := client.R().
			SetBody(map[string]any{
				"method":  "eth_getAccount",
				"params":  []string{args[0], "latest"},
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
	rootCmd.AddCommand(accountCmd)
}
