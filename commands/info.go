package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// This command shows status info for server(s)
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print information about running minecraft servers",
	Long:  `Display things like player numbers and start time`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Info coming soon")
	},
  }