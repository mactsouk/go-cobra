/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// twoCmd represents the two command
var twoCmd = &cobra.Command{
	Use:     "two",
	Aliases: []string{"cmd2"},
	Short:   "Command two short message",
	Long:    `This is the long description of command two.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("two called")
		username, _ := cmd.Flags().GetString("username")
		if username == "" {
			fmt.Println("Not a valid username:", username)
			return
		}
		fmt.Println("Hello", username)
	},
}

func init() {
	rootCmd.AddCommand(twoCmd)
	twoCmd.Flags().StringP("username", "u", "Mike", "Username value")
}
