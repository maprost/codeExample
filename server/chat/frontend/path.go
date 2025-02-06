package frontend

import (
	"runtime"
	"strings"
)

func BasePath() string {
	const thisFile = "path.go"

	_, filename, _, _ := runtime.Caller(0)
	filename = strings.TrimSuffix(filename, thisFile)
	return filename
}
