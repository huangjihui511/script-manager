package pkg

import (
	"fmt"
	"os/exec"
	"runtime"
)

func shortenSentence(rawSent string, maxLength int) string {
	if len(rawSent) < maxLength {
		return rawSent
	}
	return fmt.Sprintf("%s..", rawSent[:maxLength])
}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}
