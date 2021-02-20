/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// A Global variable
var Special = "This is a special message."

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cobra",
	Short: "A brief description of your application",
	Long:  `This is a long description of the commnad line application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("directory", "d", "/tmp", "Path to use.")
	rootCmd.PersistentFlags().Uint("depth", 2, "Depth of search.")

	viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))
	viper.BindPFlag("depth", rootCmd.PersistentFlags().Lookup("depth"))
}
