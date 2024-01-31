package main

import (
	"github.com/anotherhadi/wtui-components/notification"
	"github.com/anotherhadi/wtui-components/wtopts"
)

func main() {
	opts := wtopts.DefaultOpts()
	opts.MaxCols = 30
	notification.Notification("Success notification", "success", opts)
	notification.Notification("Warning notification", "warning", opts)
	notification.Notification("Error notification", "error", opts)

	opts.Style = "double"

	notification.Notification("Info notification", "info", opts)
	notification.Notification("Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.", "", opts)
}
