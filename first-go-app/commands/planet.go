package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "dagobah",
	Short: `Dagobah is an awesome planet`,
	Long: `Dagobah provides planet style RSS aggregation. It is inspired by
	python planet. Simple yaml config and it's own server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dagobah runs")
	},
}
