package launchitem

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func listApplicationLoginItems() ([]LaunchItem, error) {
	applications, err := os.ReadDir("/Applications")
	if err != nil {
		return nil, err
	}

	var items []LaunchItem
	for _, application := range applications {
		// Only consider directories that end with .app
		if application.IsDir() && strings.HasSuffix(application.Name(), ".app") {
			appName := application.Name()[0 : len(application.Name())-4]
			loginItems, err := os.ReadDir(
				fmt.Sprintf("/Applications/%s/Contents/Library/LoginItems", application.Name()))
			if err != nil {
				continue
			}

			for _, loginItem := range loginItems {
				// Only consider directories that end with .app
				if loginItem.IsDir() && strings.HasSuffix(loginItem.Name(), ".app") {
					loginItemName := loginItem.Name()[0 : len(loginItem.Name())-4]
					loginItemPath := fmt.Sprintf("/Applications/%s/Contents/Library/LoginItems/%s",
						application.Name(), loginItem.Name())

					fileInfo, err := os.Stat(loginItemPath)
					if err != nil {
						return nil, err
					}

					loginItemEnabled := false
					if fileInfo.Mode()&0111 != 0 {
						loginItemEnabled = true
					}

					id := sha256.New()
					id.Write([]byte(loginItemPath))

					items = append(items, LaunchItem{
						Id:      fmt.Sprintf("%x", id.Sum(nil))[0:8],
						AppName: appName,
						Name:    loginItemName,
						Path:    loginItemPath,
						Type:    Application,
						Enabled: loginItemEnabled,
					})
				}
			}
		}
	}
	return items, nil
}
