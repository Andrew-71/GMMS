package commands

import (
	"fmt"
	"os"
	"a71.su/gmms/tui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gmms",
	Short: "GMMS Lets you easily manage Minecraft servers from CLI",
	Long: `A very simple CLI tool that simplifies some parts of personal Minecraft server management. 
	Written by Andrew_7_1 in Go.
	Complete documentation is available at https://github.com/Andrew-71/GMMS`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tui.AsciiLogo())
		fmt.Println("Go Manage Minecraft Server v0.0")  // TODO: How to add versions?
	},
  }

  func init() {
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(psaCmd)
	rootCmd.AddCommand(listCommand)
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }