package input

// #cgo LDFLAGS: -lX11 -lXtst
// #include <X11/Xlib.h>
// #include <X11/extensions/XTest.h>
import "C"

var (
	screenX float32
	screenY float32
)

func init() {
	ResetScreen()
}

func ResetScreen() {
	CloseDisplay()
	display := getDisplay()
	screen := C.XDefaultScreen(display)
	screenX = float32(C.XDisplayWidth(display, screen))
	screenY = float32(C.XDisplayHeight(display, screen))
}

func mouseInput(button C.uint, isPress C.int) {
	display := getDisplay()
	C.XTestFakeButtonEvent(display, button, isPress, 0)
	C.XSync(display, 0)
}

func MouseMove(x, y int32) {
	display := getDisplay()
	C.XWarpPointer(display, 0, C.XDefaultRootWindow(display), 0, 0, 0, 0, C.int(x), C.int(y))
	C.XSync(display, 0)
}

func MouseMoveUint16(x, y uint16) {
	MouseMoveWithRatio(float32(x)/65536, float32(y)/65536)
}

func MouseMoveWithRatio(x, y float32) {
	MouseMove(int32(x*screenX), int32(y*screenY))
}

func MouseDown(button uint8) {
	switch button {
	case 0:
		mouseInput(1, 1)
	case 1:
		mouseInput(2, 1)
	case 2:
		mouseInput(3, 1)
	}
	onMouseDown(button)
}

// func MouseLeftDown() {
// 	mouseInput(1, 1)
// }

// func MouseMiddleDown() {
// 	mouseInput(2, 1)
// }

// func MouseRightDown() {
// 	mouseInput(3, 1)
// }

func mouseUp(button uint8) {
	switch button {
	case 0:
		mouseInput(1, 0)
	case 1:
		mouseInput(2, 0)
	case 2:
		mouseInput(3, 0)
	}
}

// func MouseLeftUp() {
// 	mouseInput(1, 0)
// }

// func MouseMiddleUp() {
// 	mouseInput(2, 0)
// }

// func MouseRightUp() {
// 	mouseInput(3, 0)
// }

func MouseWheel(delta int32) {
	var button C.uint = 4 // 4-up, 5-down
	magnitude := delta
	if delta < 0 {
		button = 5
		magnitude = -delta
	}
	display := getDisplay()
	for i := int32(0); i < magnitude; i++ {
		C.XTestFakeButtonEvent(display, button, 1, 0)
		C.XTestFakeButtonEvent(display, button, 0, 0)
	}
	C.XSync(display, 0)
}
