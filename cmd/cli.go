/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/joaops3/go-hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, action, productId, productName, productPrice)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "enabled", "Enable / Disable product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "PRODUCT ID")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "PRODUCT Name")
	cliCmd.Flags().Float64VarP(&productPrice, "name", "n", 0, "PRODUCT price")
}
