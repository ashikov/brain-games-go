package cmd

import (
	"github.com/ashikov/brain-games/src/games"
	"github.com/spf13/cobra"
)

var progressionCmd = &cobra.Command{
	Use:   "progression",
	Short: "You need to calculate missed arithmetic progression term",
	Long: "You need to calculate missed arithmetic progression term",
	Run: func(cmd *cobra.Command, args []string) {
		games.RunProgression()
	},
}

func init() {
	runCmd.AddCommand(progressionCmd)
}
