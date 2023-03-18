package commands

import (
	"fmt"

	"a71.su/gmms/tui"
	"a71.su/gmms/utils"
	"github.com/spf13/cobra"
)

// This command starts a server
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start minecraft server(s)",
	Long:  `Begin execution of one or more servers`,
	Run: func(cmd *cobra.Command, args []string) {
		servers := utils.GetAllServers()

		if len(args) == 0 {
			fmt.Println(tui.Red + "Needs at least one server argument to start" + tui.Reset)
		} else {
			for i := 0; i < len(args); i++ {
				found := false
				for j := 0; j < len(servers.Servers); j++ {
					if servers.Servers[j].Name == args[i] {
						if err := utils.StartServer(args[i]); err != nil {
							fmt.Println(tui.Red + "[" + args[i] + "] Error starting server: ", err, tui.Reset)
						} else {
							fmt.Println("[" + tui.Cyan + args[i] + tui.Reset + "] Started server")
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
