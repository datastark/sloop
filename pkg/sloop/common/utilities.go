package common

import (
	"fmt"
	"path"
	"strings"
)

func BoolToFloat(value bool) float64 {
	if value {
		return 1
	}
	return 0
}

func ParseKey(key string) (error, []string) {
	parts := strings.Split(key, "/")
	if len(parts) != 7 {
		return fmt.Errorf("key should have 6 parts: %v", key), parts
	}
	if parts[0] != "" {
		return fmt.Errorf("key should start with /: %v", key), parts
	}

	return nil, parts
}

func Contains(stringList []string, elem string) bool {
	for _, str := range stringList {
		if str == elem {
			return true
		}
	}
	return false
}

func GetFilePath(filePath string, fileName string) string {
	return path.Join(filePath, fileName)
}
