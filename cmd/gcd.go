package cmd

import (
	"github.com/ashikov/brain-games/src/games"
	"github.com/spf13/cobra"
)

var gcdCmd = &cobra.Command{
	Use:   "gcd",
	Short: "You'll get two numbers and have to calculate their gratest common divisor",
	Long: "You'll get two numbers and have to calculate their gratest common divisor",
	Run: func(cmd *cobra.Command, args []string) {
		games.RunGcd()
	},
}

func init() {
	runCmd.AddCommand(gcdCmd)
}
