package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"a71.su/gmms/utils"
	"a71.su/gmms/tui"
	"strconv"
)

// This command lists all added servers together with basic info

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all added minecraft servers",
	Long:  `Display all servers added into GMMS`,
	Run: func(cmd *cobra.Command, args []string) {
		servers := utils.GetAllServers()

		fmt.Println("=== Servers ===")
		fmt.Println(strconv.Itoa(len(servers.Servers)) + " added")

		for i := 0; i < len(servers.Servers); i++ {
			fmt.Println(tui.ServerFrame(servers.Servers[i]) + "\n")
		}
	},
  }