package input

// #cgo LDFLAGS: -framework Security
// #cgo LDFLAGS: -framework CoreGraphics
// #include <CoreGraphics/CoreGraphics.h>
import "C"
import (
	"time"
	"unicode/utf16"
)

var KeyboardInputUnicodeSleep = time.Millisecond

func keyboardInput(vk uint16, down bool) {
	event := C.CGEventCreateKeyboardEvent(0, C.CGKeyCode(vk), C.bool(down))
	C.CGEventPost(C.kCGSessionEventTap, event)
	C.CFRelease(C.CFTypeRef(event))
}

func KeyboardDown(vk uint16) {
	keyboardInput(vk, true)
	onKeyboardDown(vk)
}

func keyboardUp(vk uint16) {
	keyboardInput(vk, false)
}

func KeyboardInputString(v string) {
	chars := utf16.Encode([]rune(v))
	event := C.CGEventCreateKeyboardEvent(0, 0, true)
	C.CGEventKeyboardSetUnicodeString(event, C.UniCharCount(len(chars)), (*C.UniChar)(&chars[0]))
	C.CGEventPost(C.kCGSessionEventTap, event)
	C.CFRelease(C.CFTypeRef(event))
}

func CloseDisplay() {
}
