package cmd

import (
	"github.com/ashikov/brain-games/src/games"
	"github.com/spf13/cobra"
)

// primeCmd represents the prime command
var primeCmd = &cobra.Command{
	Use:   "prime",
	Short: "You need to define whether the displayed number is prime or not",
	Long: `You need to define whether the displayed number is prime or not`,
	Run: func(cmd *cobra.Command, args []string) {
		games.RunPrime()
	},
}

func init() {
	runCmd.AddCommand(primeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// primeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// primeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
