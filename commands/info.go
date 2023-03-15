package commands

import (
	"fmt"

	"a71.su/gmms/tui"
	"a71.su/gmms/utils"
	"github.com/spf13/cobra"
)

// This command shows status info for server(s)
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print information about minecraft server(s)",
	Long:  `Display more detailed report on some servers, showing more information than list command`,
	Run: func(cmd *cobra.Command, args []string) {
		servers := utils.GetAllServers()
		
		if len(args) == 0 {
			for i := 0; i < len(servers.Servers); i++ {
				fmt.Println(tui.ServerFrame(servers.Servers[i]))
			}
		} else {
			for i := 0; i < len(args); i++ {
				found := false
				for j := 0; j < len(servers.Servers); j++ {
					if servers.Servers[j].Name == args[i] {
						fmt.Println(tui.ServerFrame(servers.Servers[j]))
						found = true
						break
					}
				}
				if !found {
					fmt.Println(tui.Red + "Server not found: " + args[i] + tui.Reset)
				}
			}
		}
	},
  }