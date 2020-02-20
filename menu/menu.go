package menu

// Item represents the first level menu.
type Item string

const (
	// Enable is the menu item that re-enables f.lux
	Enable Item = "f.lux is off"

	// Color is the menu item to change effects.
	Color Item = "Color Effects"

	// Disable is the menu item to disable Flux.
	Disable Item = "Disable"
)
