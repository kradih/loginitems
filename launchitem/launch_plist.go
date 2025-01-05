package launchitem

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"howett.net/plist"
)

func PrintPlist() {
	data, err := os.ReadFile("/Library/LaunchAgents/com.microsoft.OneDriveStandaloneUpdater.plist")
	if err == nil {
		var parsedData map[string]interface{} = make(map[string]interface{})
		formatType, _ := plist.Unmarshal(data, parsedData)
		fmt.Println(formatType)
		fmt.Println(parsedData)
	}
}

func launchItemFromPlist(path string) (*LaunchItem, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var parsedContent map[string]interface{} = make(map[string]interface{})
	plistFormat, err := plist.Unmarshal(content, &parsedContent)
	if err != nil {
		return nil, err
	}
	if plistFormat == plist.InvalidFormat {
		return nil, fmt.Errorf("unsupported plist format")
	}

	launchItem := &LaunchItem{}

	id := sha256.New()
	id.Write([]byte(path))

	launchItem.Id = fmt.Sprintf("%x", id.Sum(nil))[0:8]
	launchItem.AppName = filepath.Base(parsedContent["Program"].(string))
	launchItem.Name = parsedContent["Label"].(string)
	launchItem.Path = path
	launchItem.Enabled = true

	return launchItem, nil
}

func listLaunchItems(dirPath string, _type int) ([]LaunchItem, error) {
	var launchItems []LaunchItem = make([]LaunchItem, 0)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.Type().IsRegular() && strings.HasSuffix(entry.Name(), ".plist") {
			launchItem, err := launchItemFromPlist(fmt.Sprintf("%s/%s", dirPath, entry.Name()))
			if err != nil {
				continue
			}
			launchItem.Type = _type
			launchItems = append(launchItems, *launchItem)
		}
	}

	return launchItems, nil
}
