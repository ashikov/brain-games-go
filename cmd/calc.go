/*
Copyright © 2022 Roman Ashikov pickle@ashikov.ru
*/
package cmd

import (
	// "fmt"
	"github.com/ashikov/brain-games/helpers"
	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "Adding, subtraction or multiplication of two values",
	Long: `Adding, subtraction or multiplication of two values`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("calc called")
		helpers.Engine()
	},
}

func init() {
	runCmd.AddCommand(calcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}