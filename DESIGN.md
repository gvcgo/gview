This document lists architectural details of gview.

# Focus-related style attributes are unset by default

This applies to all widgets except Button and TabbedPanels, which require a
style change to indicate focus. See [ColorUnset](https://pkg.go.dev/github.com/gvcgo/gview#pkg-variables).

# Widgets always use `sync.RWMutex`

See [#30](https://github.com/gvcgo/gview/issues/30).
