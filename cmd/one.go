/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// oneCmd represents the one command
var oneCmd = &cobra.Command{
	Use:     "one",
	Aliases: []string{"cmd1"},
	Short:   "Command one",
	Long:    `This is the description for one.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("one called")

		dir := viper.GetString("directory")
		if dir == "" {
			fmt.Println("Dir is", dir)
		}

		depth := viper.GetUint("depth")
		fmt.Println(dir, depth)
	},
}

func init() {
	rootCmd.AddCommand(oneCmd)
}
