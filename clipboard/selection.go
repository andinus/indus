package clipboard

import (
	"os/exec"
)

// GetSel returns the primary clipboard selection using the xclip tool
// & also returns an error (if exists) along with the primary
// clipboard selection.
func GetSel() (string, error) {
	out, err := exec.Command("xclip", "-out").Output()
	if err != nil {
		return "", err
	}

	sel := string(out)
	return sel, err
}
