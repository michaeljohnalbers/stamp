package cmd

import (
	"github.com/spf13/cobra"
)

// TODO:
//   make sure verbiage is consistent in terms of input thing and generated output thing
//   add option for program config
//      - left/right template delimiters
//      - per file/glob config (so you can switch delimiters, etc based on path/type/etc.)
//      - file/glob to exclude

var rootCmd = &cobra.Command{
	Use: "stamp",
	Short: "Stamp is a tool for generating projects from templates",
	Long: `Stamp is a CLI used to take user-defined project templates add user-defined specifics and create a project shell.`,
}

func Execute() error{
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(genCommand)
	rootCmd.AddCommand(versionCommand)
}
