package script

// https://forum.justgetflux.com/topic/1928/how-to-create-keyboard-shortcuts

// Flux is the AppleScript template to execute menu option changes.
var Flux = `
property mainItem : "%s"
-- set to "Preferences...", "Color Effects", "Disable", etc.,
-- make sure to use quote marks.

property subItem : "%s"
-- set to submenu item, if there is one. Use "for this app" with Disable,
-- to toggle disable for the current application.

if mainItem is "Disable" and subItem is "for this app" then set subItem to 3

do shell script "open -g /Applications/Flux.app"

ignoring application responses
    tell application "System Events" to tell process "Flux"
        try
            click menu bar item 1 of menu bar 2
        end try
    end tell
end ignoring

do shell script "killall 'System Events'"

tell application "System Events" to tell process "Flux"
    tell menu 1 of menu bar item 1 of menu bar 2
        click menu item mainItem
        tell menu item mainItem
            if menu 1 exists then
                click menu item subItem of menu 1
            end if
        end tell
    end tell
    if mainItem is "Preferences..." then set frontmost to true
end tell

`
