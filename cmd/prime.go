package cmd

import (
	"github.com/ashikov/brain-games/src/games"
	"github.com/spf13/cobra"
)

var primeCmd = &cobra.Command{
	Use:   "prime",
	Short: "You need to define whether the displayed number is prime or not",
	Long: "You need to define whether the displayed number is prime or not",
	Run: func(cmd *cobra.Command, args []string) {
		games.RunPrime()
	},
}

func init() {
	runCmd.AddCommand(primeCmd)
}
