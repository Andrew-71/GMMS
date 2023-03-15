package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"a71.su/gmms/utils"
	"strconv"
)

// This command lists all added server names
var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all added minecraft servers",
	Long:  `Display list of all server names added into GMMS`,
	Run: func(cmd *cobra.Command, args []string) {
		servers := utils.GetAllServers()
		fmt.Println("=== Servers ===")
		fmt.Println(strconv.Itoa(len(servers.Servers)) + " servers added")
		for i := 0; i < len(servers.Servers); i++ {
			fmt.Println(servers.Servers[i].Name)
		}
	},
  }