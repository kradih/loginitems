package launchitem

func listSystemLaunchDaemons() ([]LaunchItem, error) {
	return listLaunchItems("/Library/LaunchDaemons", SystemLaunchDaemon)
}
