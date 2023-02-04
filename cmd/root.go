package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// root
//	| validator --status==jailed|unjailed|...
//		| votes
//		| node-info
//	| network
//		| blocktime
//		| status
//		| nodes
//		| relayers
//	| market
//		| uatom

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mintscan-alerter",
	Short: "Command line tool to alert on Cosmos Hub validator status changes",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mintscan-alerter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
