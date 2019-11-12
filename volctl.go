package volctl

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var (
	volumeRegex = regexp.MustCompile(`.*: .*\[(\d{1,3})%`)
)

// FindVolume grabs the volume percentage from the amixer get command
// output. Example output:
// Simple mixer control 'PCM',0
//   Capabilities: pvolume pvolume-joined pswitch pswitch-joined
//   Playback channels: Mono
//   Limits: Playback -10239 - 400
//   Mono: Playback 389 [100%] [3.89dB] [on]
func FindVolume(out []byte) string {
	lines := bytes.Split(out, []byte("\n"))

	for _, line := range lines {
		if !bytes.Contains(line, []byte("[")) {
			continue
		}
		m := volumeRegex.FindSubmatch(line)
		if len(m) <= 0 {
			continue
		}
		return string(m[1]) // return the first group
	}

	return ""
}

// CurrentVolume inspects amixer output to get the current volume.
func CurrentVolume() (string, error) {
	c := exec.Command("amixer", "get", "PCM")
	out, err := c.Output()
	if err != nil {
		return "", err
	}

	vol := FindVolume(out)
	if vol == "" {
		return "", fmt.Errorf("No Volume found!")
	}

	return vol, nil
}

// UpdateVolume updates the volume according to a percentage using the
// amixer command.
func UpdateVolume(vol string) error {
	if !strings.HasSuffix(vol, "%") {
		vol = fmt.Sprintf("%s%%", vol)
	}

	c := exec.Command("amixer", "set", "PCM", "--", vol)
	out, err := c.Output()
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("Failed to update the volume: %s", string(out)))
	}
	return nil
}
