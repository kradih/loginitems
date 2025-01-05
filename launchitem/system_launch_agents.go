package launchitem

func listSystemLaunchAgents() ([]LaunchItem, error) {
	return listLaunchItems("/Library/LaunchAgents", SystemLaunchAgent)
}
