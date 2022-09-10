package cmd

import (
	"github.com/ashikov/brain-games/src/games"
	"github.com/spf13/cobra"
)

// progressionCmd represents the progression command
var progressionCmd = &cobra.Command{
	Use:   "progression",
	Short: "You need to calculate missed arithmetic progression term",
	Long: `You need to calculate missed arithmetic progression term`,
	Run: func(cmd *cobra.Command, args []string) {
		games.RunProgression()
	},
}

func init() {
	runCmd.AddCommand(progressionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// progressionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// progressionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
