package input

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

var gDisplay *C.Display

func getDisplay() *C.Display {
	if gDisplay == nil {
		gDisplay = C.XOpenDisplay(nil)
	}
	return gDisplay
}

func CloseDisplay() {
	if gDisplay != nil {
		C.XCloseDisplay(gDisplay)
		gDisplay = nil
	}
}
