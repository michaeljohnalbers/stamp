package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use: "version",
	Short: "Print version information",
	Run: version,
}

func version(_ *cobra.Command, _ []string) {
	// TODO: something better than this
	fmt.Println("0.1.0")
}
