package launchitem

const (
	Application = iota
	SystemLaunchDaemon
	SystemLaunchAgent
	UserLaunchDaemon
	UserLaunchAgent
)

type LaunchItem struct {
	Id      string
	AppName string
	Name    string
	Enabled bool
	Path    string
	Type    int
}

func List() ([]LaunchItem, error) {
	applicationLoginItems, err := listApplicationLoginItems()
	if err != nil {
		return nil, err
	}

	systemLaunchDaemons, err := listSystemLaunchDaemons()
	if err != nil {
		return nil, err
	}

	systemLaunchAgents, err := listSystemLaunchAgents()
	if err != nil {
		return nil, err
	}

	userLaunchDaemons, err := listUserLaunchDaemons()
	if err != nil {
		return nil, err
	}

	userLaunchAgents, err := listUserLaunchAgents()
	if err != nil {
		return nil, err
	}

	items := append(applicationLoginItems, systemLaunchDaemons...)
	items = append(items, systemLaunchAgents...)
	items = append(items, userLaunchDaemons...)
	items = append(items, userLaunchAgents...)

	return items, nil
}
