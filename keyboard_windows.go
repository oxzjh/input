package input

import (
	"time"
	"unicode/utf16"
	"unsafe"
)

var KeyboardInputUnicodeSleep = time.Millisecond

var keyboardInputSize = int32(unsafe.Sizeof(KEYBD_INPUT{}))

func keyboardInput(vk, scan uint16, flags uint32) {
	gUser32.SendInput(1, unsafe.Pointer(&KEYBD_INPUT{
		Type: INPUT_KEYBOARD,
		Ki: KEYBDINPUT{
			WVk:     vk,
			WScan:   scan,
			DwFlags: flags,
		},
	}), keyboardInputSize)
}

func KeyboardDown(vk uint16) {
	keyboardInput(vk, 0, 0)
	onKeyboardDown(vk)
}

func keyboardUp(vk uint16) {
	keyboardInput(vk, 0, KEYEVENTF_KEYUP)
}

func KeyboardInputString(v string) {
	unicodes := utf16.Encode([]rune(v))
	for _, unicode := range unicodes {
		keyboardInput(0, unicode, KEYEVENTF_UNICODE)
		time.Sleep(KeyboardInputUnicodeSleep)
	}
}

func CloseDisplay() {
}
