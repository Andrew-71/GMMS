package tui

import (
	"os"
	"runtime"
	"strings"
	"time"

	"a71.su/gmms/utils"
	"github.com/briandowns/spinner"
)

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

func init() {
	// No colours on windows :/
	if runtime.GOOS == "windows" {
		Reset  = ""
		Red    = ""
		Green  = ""
		Yellow = ""
		Blue   = ""
		Purple = ""
		Cyan   = ""
		Gray   = ""
		White  = ""
	}
}


func AsciiLogo() string {
	return Cyan + `
                ▇▇▇▇▇▇  ▇▇▇    ▇▇▇ ▇▇▇    ▇▇▇ ▇▇▇▇▇▇▇ 
      ▇▇▇▇▇▇▇▇ ▇▇       ▇▇▇▇  ▇▇▇▇ ▇▇▇▇  ▇▇▇▇ ▇▇      
▇▇▇▇▇▇▇▇▇▇▇▇▇▇ ▇▇   ▇▇▇ ▇▇ ▇▇▇▇ ▇▇ ▇▇ ▇▇▇▇ ▇▇ ▇▇▇▇▇▇▇ 
        ▇▇▇▇▇▇ ▇▇    ▇▇ ▇▇  ▇▇  ▇▇ ▇▇  ▇▇  ▇▇      ▇▇ 
                ▇▇▇▇▇▇  ▇▇      ▇▇ ▇▇      ▇▇ ▇▇▇▇▇▇▇ 
	` + Reset
}

func RightPad(str, ch string, length int) string {
	return str + strings.Repeat(ch, (utils.Max(0, length - len(str))))
}

// ooo left-pad reference???? NPM????
func LeftPad(str, ch string, length int) string {
	return strings.Repeat(ch, (utils.Max(0, length - len(str)))) + str
}

// Easy way to make default spinner
func GetSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
	s.Color("cyan")
	return s
}

func ServerFrame(server utils.Server) string {
	is_running := false
	processes := utils.GetAllProcesses()
	for i := 0; i < len(processes.Processes); i++ {
		if processes.Processes[i].Name == server.Name {
			is_running = true
			break
		}
	}

	out := "[" + Cyan + server.Name + Reset + "]\n"
	out += "\tRunning: "
	if is_running {
		out += Green + "Yes\n"
	} else {
		out += Red + "No\n"
	}
	out += Reset + "\tAdd date: " + time.Since(time.Unix(server.AddedTime, 0)).Round(time.Second).String() + " ago\n"
	out += "\tLast edited: " + time.Since(time.Unix(server.EditTime, 0)).Round(time.Second).String() + " ago\n"
	out += "\tDiscord: "
	if server.Discord.Enabled {
		out += Green + "Yes\n"
	} else {
		out += Red + "No\n"
	}
	out += Reset + "\tJar: " + server.JarDir
	return out

	/*
	Server name --------- < to 20
	Running: no
	Add date: March 1st 1984
	Edit date: June 5th 2015
	Jar: /main_ssd/4893/sja
	*/
}