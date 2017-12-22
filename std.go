package libsys

import (
	"os"

	"github.com/pkg/errors"
)

// IsPiped is a convenience function for checking if there is content on the other
// end of stdin
func IsPiped() (bool, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, errors.Wrap(err, "Failed to stat stdin")
	}
	if fi.Mode()&os.ModeCharDevice == 0 {
		return true, nil
	}
	return false, nil
}
