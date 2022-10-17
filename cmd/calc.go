package cmd

import (
	"github.com/ashikov/brain-games/internal/games"
	"github.com/spf13/cobra"
)

var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "Adding, subtraction or multiplication of two values",
	Long:  "Adding, subtraction or multiplication of two values",
	Run: func(cmd *cobra.Command, args []string) {
		games.RunCalc()
	},
}

func init() {
	runCmd.AddCommand(calcCmd)
}
