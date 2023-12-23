/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// txCmd represents the tx command
var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("lack of tx hash")
		}

		rpcUrl := viper.GetString("rpc-url")
		client := req.C()
		resp, err := client.R().
			SetBody(map[string]any{
				"method":  "eth_getTransactionByHash",
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
	rootCmd.AddCommand(txCmd)

	txCmd.Flags().StringP("rpc-url", "r", "https://eth.llamarpc.com", "Set rpc url")

	if err := viper.BindPFlags(txCmd.Flags()); err != nil {
		panic(err)
	}
}
