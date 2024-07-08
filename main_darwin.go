//go:build darwin

package main

import (
	"fmt"
	"os"

	"github.com/kradih/loginitems/launchitem"

	"github.com/olekukonko/tablewriter"
)

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "help":
		help()
	case "list":
		list()
	}
}

func help() {
	fmt.Println("Usage: loginitems <help|list|enable <id>|disable <id>>")
}

func list() error {
	items, err := launchitem.List()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"ID", "App Name", "Name", "Enabled", "Path", "Type"})

	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_CENTER,  // ID
		tablewriter.ALIGN_CENTER,  // App Name
		tablewriter.ALIGN_CENTER,  // Name
		tablewriter.ALIGN_CENTER,  // Enabled
		tablewriter.ALIGN_DEFAULT, // Path
		tablewriter.ALIGN_CENTER,  // Type
	})

	table.SetAutoWrapText(false)

	for _, item := range items {
		enabled := ""
		if item.Enabled {
			enabled = "x"
		}

		var _type string
		switch item.Type {
		case launchitem.Application:
			_type = "Application"
		case launchitem.SystemLaunchDaemon:
			_type = "SystemLaunchDaemon"
		case launchitem.SystemLaunchAgent:
			_type = "SystemLaunchAgent"
		case launchitem.UserLaunchDaemon:
			_type = "UserLaunchDaemon"
		case launchitem.UserLaunchAgent:
			_type = "UserLaunchAgent"
		default:
			_type = "Unknown"
		}

		table.Append([]string{item.Id, item.AppName, item.Name, enabled, item.Path, _type})
	}
	table.Render()

	return nil
}
