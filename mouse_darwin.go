package input

// #cgo LDFLAGS: -framework Security
// #cgo LDFLAGS: -framework CoreGraphics
// #cgo LDFLAGS: -framework ApplicationServices
// #include <CoreGraphics/CoreGraphics.h>
// #include <ApplicationServices/ApplicationServices.h>
import "C"

var (
	screenX float32
	screenY float32
)

func init() {
	ResetScreen()
}

func ResetScreen() {
	displayID := C.CGMainDisplayID()
	screenX = float32(C.CGDisplayPixelsWide(displayID))
	screenY = float32(C.CGDisplayPixelsHigh(displayID))
}

func mousePosition() C.CGPoint {
	event := C.CGEventCreate(0)
	defer C.CFRelease(C.CFTypeRef(event))
	return C.CGEventGetLocation(event)
}

func mouseInput(typ C.CGEventType, point C.CGPoint, button C.CGMouseButton) {
	event := C.CGEventCreateMouseEvent(0, typ, point, button)
	C.CGEventPost(C.kCGSessionEventTap, event)
	C.CFRelease(C.CFTypeRef(event))
}

func MouseMove(x, y int32) {
	mouseInput(C.kCGEventMouseMoved, C.CGPointMake(C.CGFloat(x), C.CGFloat(y)), C.kCGMouseButtonLeft)
}

func MouseMoveUint16(x, y uint16) {
	MouseMoveWithRatio(float32(x)/65536, float32(y)/65536)
}

func MouseMoveWithRatio(x, y float32) {
	MouseMove(int32(screenX*x), int32(screenY*y))
}

func MouseDown(button uint8) {
	switch button {
	case 0:
		mouseInput(C.kCGEventLeftMouseDown, mousePosition(), C.kCGMouseButtonLeft)
	case 1:
		mouseInput(C.kCGEventOtherMouseDown, mousePosition(), C.kCGMouseButtonCenter)
	case 2:
		mouseInput(C.kCGEventRightMouseDown, mousePosition(), C.kCGMouseButtonRight)
	}
	onMouseDown(button)
}

// func MouseLeftDown() {
// 	mouseInput(C.kCGEventLeftMouseDown, mousePosition(), C.kCGMouseButtonLeft)
// }

// func MouseMiddleDown() {
// 	mouseInput(C.kCGEventOtherMouseDown, mousePosition(), C.kCGMouseButtonCenter)
// }

// func MouseRightDown() {
// 	mouseInput(C.kCGEventRightMouseDown, mousePosition(), C.kCGMouseButtonRight)
// }

func mouseUp(button uint8) {
	switch button {
	case 0:
		mouseInput(C.kCGEventLeftMouseUp, mousePosition(), C.kCGMouseButtonLeft)
	case 1:
		mouseInput(C.kCGEventOtherMouseUp, mousePosition(), C.kCGMouseButtonCenter)
	case 2:
		mouseInput(C.kCGEventRightMouseUp, mousePosition(), C.kCGMouseButtonRight)
	}
}

// func MouseLeftUp() {
// 	mouseInput(C.kCGEventLeftMouseUp, mousePosition(), C.kCGMouseButtonLeft)
// }

// func MouseMiddleUp() {
// 	mouseInput(C.kCGEventOtherMouseUp, mousePosition(), C.kCGMouseButtonCenter)
// }

// func MouseRightUp() {
// 	mouseInput(C.kCGEventRightMouseUp, mousePosition(), C.kCGMouseButtonRight)
// }

func MouseWheel(delta int32) {
	event := C.CGEventCreateScrollWheelEvent2(0, C.kCGScrollEventUnitPixel, 1, C.int32_t(delta), 0, 0)
	C.CGEventPost(C.kCGHIDEventTap, event)
	C.CFRelease(C.CFTypeRef(event))
}
