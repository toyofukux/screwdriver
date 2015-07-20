package utils

import (
	"os"
	"strings"
)

const (
	// EnvPrefix is the prefix to load environment variable.
	EnvPrefix = "SCREW_"
)

// LoadScrewEnvs is laod environment variable and return list of environment.
func LoadScrewEnvs() map[string]string {
	data := os.Environ()
	items := make(map[string]string)
	for _, val := range data {
		splits := strings.SplitN(val, "=", 2)
		key := splits[0]
		if strings.HasPrefix(key, EnvPrefix) {
			items[key] = splits[1]
		}
	}
	return items
}
