package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// This command shows status info for server(s)

var psaCmd = &cobra.Command{
	Use:   "psa",
	Short: "Send a message to server player",
	Long:  `Print a message to console (or elsewhere) for all players on one or several servers`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("PSA coming soon")
	},
  }