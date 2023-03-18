package commands

import (
	"fmt"

	"a71.su/gmms/tui"
	"a71.su/gmms/utils"
	"github.com/spf13/cobra"
)

// This command stops a server
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop minecraft server(s)",
	Long:  `Stop execution of one or more servers`,
	Run: func(cmd *cobra.Command, args []string) {
		servers := utils.GetAllServers()

		if len(args) == 0 {
			fmt.Println(tui.Red + "Needs at least one server argument to stop" + tui.Reset)
		} else {
			for i := 0; i < len(args); i++ {
				found := false
				for j := 0; j < len(servers.Servers); j++ {
					if servers.Servers[j].Name == args[i] {
						if err := utils.StopServer(args[i]); err != nil {
							fmt.Println(tui.Red + "[" + args[i] + "] Error stopping server: ", err, tui.Reset)
						} else {
							fmt.Println("[" + tui.Cyan + args[i] + tui.Reset + "] stopped server")
						}
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
