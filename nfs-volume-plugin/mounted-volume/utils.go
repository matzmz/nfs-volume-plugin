package mountedvolume

import (
	"fmt"
	"os"
	"strings"
)

const (
	ShareSplitIndentifer = "#"
)

func createDest(dest string) error {
	fi, err := os.Lstat(dest)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if fi != nil && !fi.IsDir() {
		return fmt.Errorf("%v already exist and it's not a directory", dest)
	}
	return nil
}

func getBooleanOption(opts map[string]string, key string) bool {
	if option, exists := opts[key]; exists {
		if (strings.ToLower(option) == "true") {
			return true
		}
	}
	return false
}
