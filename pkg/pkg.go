package pkg

import (
	"fmt"
	"os/exec"

	"github.com/DavideFauri/fluxflow/menu"
	"github.com/DavideFauri/fluxflow/menu/submenu"
	"github.com/DavideFauri/fluxflow/script"
)

// Run executes the AppleScript with the chosen options.
func Run(primary menu.Item, secondary submenu.Item) error {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(script.Flux, primary, secondary))

	return cmd.Run()
}
