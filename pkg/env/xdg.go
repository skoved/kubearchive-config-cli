// Copyright skoved

package env

import (
	"fmt"

	"github.com/adrg/xdg"
)

const subDir = "kac"

func XdgDataHome() string {
	return fmt.Sprintf("%s/%s", xdg.DataHome, subDir)
}
