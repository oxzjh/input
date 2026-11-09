package input

var keymap = map[string]uint16{
	"Digit0": 0x30,
	"Digit1": 0x31,
	"Digit2": 0x32,
	"Digit3": 0x33,
	"Digit4": 0x34,
	"Digit5": 0x35,
	"Digit6": 0x36,
	"Digit7": 0x37,
	"Digit8": 0x38,
	"Digit9": 0x39,

	"KeyA": 0x41,
	"KeyB": 0x42,
	"KeyC": 0x43,
	"KeyD": 0x44,
	"KeyE": 0x45,
	"KeyF": 0x46,
	"KeyG": 0x47,
	"KeyH": 0x48,
	"KeyI": 0x49,
	"KeyJ": 0x4A,
	"KeyK": 0x4B,
	"KeyL": 0x4C,
	"KeyM": 0x4D,
	"KeyN": 0x4E,
	"KeyO": 0x4F,
	"KeyP": 0x50,
	"KeyQ": 0x51,
	"KeyR": 0x52,
	"KeyS": 0x53,
	"KeyT": 0x54,
	"KeyU": 0x55,
	"KeyV": 0x56,
	"KeyW": 0x57,
	"KeyX": 0x58,
	"KeyY": 0x59,
	"KeyZ": 0x5A,

	"F1":  0x70, // VK_F1
	"F2":  0x71, // VK_F2
	"F3":  0x72, // VK_F3
	"F4":  0x73, // VK_F4
	"F5":  0x74, // VK_F5
	"F6":  0x75, // VK_F6
	"F7":  0x76, // VK_F7
	"F8":  0x77, // VK_F8
	"F9":  0x78, // VK_F9
	"F10": 0x79, // VK_F10
	"F11": 0x7A, // VK_F11
	"F12": 0x7B, // VK_F12

	"Insert":   0x2D, // VK_INSERT
	"Home":     0x24, // VK_HOME
	"PageUp":   0x21, // VK_PRIOR
	"Delete":   0x2E, // VK_DELETE
	"End":      0x23, // VK_END
	"PageDown": 0x22, // VK_NEXT

	"ArrowLeft":  0x25, // VK_LEFT
	"ArrowUp":    0x26, //VK_UP
	"ArrowRight": 0x27, //VK_RIGHT
	"ArrowDown":  0x28, //VK_DOWN

	"Numpad0": 0x60, // VK_NUMPAD0
	"Numpad1": 0x61, // VK_NUMPAD1
	"Numpad2": 0x62, // VK_NUMPAD2
	"Numpad3": 0x63, // VK_NUMPAD3
	"Numpad4": 0x64, // VK_NUMPAD4
	"Numpad5": 0x65, // VK_NUMPAD5
	"Numpad6": 0x66, // VK_NUMPAD6
	"Numpad7": 0x67, // VK_NUMPAD7
	"Numpad8": 0x68, // VK_NUMPAD8
	"Numpad9": 0x69, // VK_NUMPAD9

	"NumLock":        0x90, // VK_NUMLOCK
	"NumpadDivide":   0x6F, // VK_DIVIDE
	"NumpadMultiply": 0x6A, // VK_MULTIPLY
	"NumpadSubtract": 0x6D, // VK_SUBTRACT
	"NumpadAdd":      0x6B, // VK_ADD
	"NumpadDecimal":  0x6E, // VK_DECIMAL
	"NumpadEnter":    0x0D, // VK_RETURN

	"Escape":       0x1B, // VK_ESCAPE
	"Backquote":    0xC0, // VK_OEM_3
	"Minus":        0xBD, // VK_OEM_MINUS
	"Equal":        0xBB, // VK_OEM_PLUS
	"Backspace":    0x08, // VK_BACK
	"Tab":          0x09, // VK_TAB
	"BracketLeft":  0xDB, // VK_OEM_4
	"BracketRight": 0xDD, // VK_OEM_6
	"Backslash":    0xDC, // VK_OEM_5
	"CapsLock":     0x14, // VK_CAPITAL
	"Semicolon":    0xBA, // VK_OEM_1
	"Quote":        0xDE, // VK_OEM_7
	"Enter":        0x0D, // VK_RETURN
	"ShiftLeft":    0xA0, // VK_LSHIFT
	"Comma":        0xBC, // VK_OEM_COMMA
	"Period":       0xBE, // VK_OEM_PERIOD
	"Slash":        0xBF, // VK_OEM_2
	"ShiftRight":   0xA1, // VK_RSHIFT
	"ControlLeft":  0xA2, // VK_LCONTROL
	"MetaLeft":     0x5B, // VK_LWIN
	"AltLeft":      0xA4, // VK_LMENU
	"Space":        0x20, // VK_SPACE
	"AltRight":     0xA5, // VK_RMENU
	"ContextMenu":  0x5D, // VK_APPS
	"ControlRight": 0xA3, // VK_RCONTROL
}
