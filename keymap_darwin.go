package input

var keymap = map[string]uint16{
	"Digit0": 0x1D, // kVK_ANSI_0
	"Digit1": 0x12, // kVK_ANSI_1
	"Digit2": 0x13, // kVK_ANSI_2
	"Digit3": 0x14, // kVK_ANSI_3
	"Digit4": 0x15, // kVK_ANSI_4
	"Digit5": 0x17, // kVK_ANSI_5
	"Digit6": 0x16, // kVK_ANSI_6
	"Digit7": 0x1A, // kVK_ANSI_7
	"Digit8": 0x1C, // kVK_ANSI_8
	"Digit9": 0x19, // kVK_ANSI_9

	"KeyA": 0x00, // kVK_ANSI_A
	"KeyB": 0x0B, // kVK_ANSI_B
	"KeyC": 0x08, // kVK_ANSI_C
	"KeyD": 0x02, // kVK_ANSI_D
	"KeyE": 0x0E, // kVK_ANSI_E
	"KeyF": 0x03, // kVK_ANSI_F
	"KeyG": 0x05, // kVK_ANSI_G
	"KeyH": 0x04, // kVK_ANSI_H
	"KeyI": 0x22, // kVK_ANSI_I
	"KeyJ": 0x26, // kVK_ANSI_J
	"KeyK": 0x28, // kVK_ANSI_K
	"KeyL": 0x25, // kVK_ANSI_L
	"KeyM": 0x2E, // kVK_ANSI_M
	"KeyN": 0x2D, // kVK_ANSI_N
	"KeyO": 0x1F, // kVK_ANSI_O
	"KeyP": 0x23, // kVK_ANSI_P
	"KeyQ": 0x0C, // kVK_ANSI_Q
	"KeyR": 0x0F, // kVK_ANSI_R
	"KeyS": 0x01, // kVK_ANSI_S
	"KeyT": 0x11, // kVK_ANSI_T
	"KeyU": 0x20, // kVK_ANSI_U
	"KeyV": 0x09, // kVK_ANSI_V
	"KeyW": 0x0D, // kVK_ANSI_W
	"KeyX": 0x07, // kVK_ANSI_X
	"KeyY": 0x10, // kVK_ANSI_Y
	"KeyZ": 0x06, // kVK_ANSI_Z

	"F1":  0x7A, // kVK_F1
	"F2":  0x78, // kVK_F2
	"F3":  0x63, // kVK_F3
	"F4":  0x76, // kVK_F4
	"F5":  0x60, // kVK_F5
	"F6":  0x61, // kVK_F6
	"F7":  0x62, // kVK_F7
	"F8":  0x64, // kVK_F8
	"F9":  0x65, // kVK_F9
	"F10": 0x6D, // kVK_F10
	"F11": 0x67, // kVK_F11
	"F12": 0x6F, // kVK_F12

	"Home":     0x73, // kVK_Home
	"PageUp":   0x74, // kVK_PageUp
	"Delete":   0x75, // kVK_ForwardDelete
	"End":      0x77, // kVK_End
	"PageDown": 0x79, // kVK_PageDown

	"ArrowLeft":  0x7B, // kVK_LeftArrow
	"ArrowUp":    0x7E, // kVK_UpArrow
	"ArrowRight": 0x7C, // kVK_RightArrow
	"ArrowDown":  0x7D, // kVK_DownArrow

	"Numpad0": 0x52, // kVK_ANSI_Keypad0
	"Numpad1": 0x53, // kVK_ANSI_Keypad1
	"Numpad2": 0x54, // kVK_ANSI_Keypad2
	"Numpad3": 0x55, // kVK_ANSI_Keypad3
	"Numpad4": 0x56, // kVK_ANSI_Keypad4
	"Numpad5": 0x57, // kVK_ANSI_Keypad5
	"Numpad6": 0x58, // kVK_ANSI_Keypad6
	"Numpad7": 0x59, // kVK_ANSI_Keypad7
	"Numpad8": 0x5B, // kVK_ANSI_Keypad8
	"Numpad9": 0x5C, // kVK_ANSI_Keypad9

	"NumpadDivide":   0x4B, // kVK_ANSI_KeypadDivide
	"NumpadMultiply": 0x43, // kVK_ANSI_KeypadMultiply
	"NumpadSubtract": 0x4E, // kVK_ANSI_KeypadMinus
	"NumpadAdd":      0x45, // kVK_ANSI_KeypadPlus
	"NumpadDecimal":  0x41, // kVK_ANSI_KeypadDecimal
	"NumpadEnter":    0x4C, // kVK_ANSI_KeypadEnter

	"Escape":       0x35, // kVK_Escape
	"Backquote":    0x32, // kVK_ANSI_Grave
	"Minus":        0x1B, // kVK_ANSI_Minus
	"Equal":        0x18, // kVK_ANSI_Equal
	"Backspace":    0x33, // kVK_Delete
	"Tab":          0x30, // kVK_Tab
	"BracketLeft":  0x21, // kVK_ANSI_LeftBracket
	"BracketRight": 0x1E, // kVK_ANSI_RightBracket
	"Backslash":    0x2A, // kVK_ANSI_Backslash
	"CapsLock":     0x39, // kVK_CapsLock
	"Semicolon":    0x29, // kVK_ANSI_Semicolon
	"Quote":        0x27, // kVK_ANSI_Quote
	"Enter":        0x24, // kVK_Return
	"ShiftLeft":    0x38, // kVK_Shift
	"Comma":        0x2B, // kVK_ANSI_Comma
	"Period":       0x2F, // kVK_ANSI_Period
	"Slash":        0x2C, // kVK_ANSI_Slash
	"ShiftRight":   0x3C, // kVK_RightShift
	"ControlLeft":  0x3B, // kVK_Control
	"MetaLeft":     0x37, // kVK_Command
	"AltLeft":      0x3A, // kVK_Option
	"Space":        0x31, // kVK_Space
	"AltRight":     0x3D, // kVK_RightOption
	"ControlRight": 0x3E, // kVK_RightControl
}
