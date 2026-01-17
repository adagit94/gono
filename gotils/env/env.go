package env

import (
	"os"
	"strings"
)

func GetEnvVarsCli() map[string]string {
	args := os.Environ()
	vars := make(map[string]string, len(args))

	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)
		vars[s[0]] = s[1]
	}

	return vars
}
