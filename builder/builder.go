package builder

import (
	"os/exec"
	"strings"
)

// RunCmd runs a build command
func RunCmd(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// ParseCmd splits the command string into cmd and arguments
func ParseCmd(cmd string) (string, []string) {
	parts := strings.Split(cmd, " ")

	if len(parts) == 1 {
		return parts[0], []string{}
	}

	return parts[0], parts[:len(parts)-1]
}
