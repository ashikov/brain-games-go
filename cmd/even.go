package cmd

import (
	"github.com/ashikov/brain-games/src/games"
	"github.com/spf13/cobra"
)

var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "You need to answer whether a given number is even or not",
	Long:  "You need to answer whether a given number is even or not",
	Run: func(cmd *cobra.Command, args []string) {
		games.RunEven()
	},
}

func init() {
	runCmd.AddCommand(evenCmd)
}
