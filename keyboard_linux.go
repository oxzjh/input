package input

/*
#cgo LDFLAGS: -lX11 -lXtst
#include <stdlib.h>
#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/XKBlib.h>
#include <X11/extensions/XTest.h>

void keyboardInputUnicode(Display *display, const char *u)
{
	KeySym keysym = XStringToKeysym(u);
	KeyCode keycode = XKeysymToKeycode(display, keysym);
	int shift = 0;

	if (keycode == 0)
	{
		static int mod = 0;
		int min, max, num;
		XDisplayKeycodes(display, &min, &max);
		KeySym *keysyms = XGetKeyboardMapping(display, min, max - min + 1, &num);
		mod = (mod + 1) & 0x7;
		keysyms[(max - min - mod - 1) * num] = keysym;
		XChangeKeyboardMapping(display, min, num, keysyms, max - min);
		XFree(keysyms);
		XSync(display, False);
		keycode = max - mod - 1;
	}
	else
		shift = XkbKeycodeToKeysym(display, keycode, 0, 0) != keysym;

	if (shift)
		XTestFakeKeyEvent(display, XKeysymToKeycode(display, XK_Shift_L), True, 0);
	XTestFakeKeyEvent(display, keycode, True, 0);
	XTestFakeKeyEvent(display, keycode, False, 0);
	if (shift)
		XTestFakeKeyEvent(display, XKeysymToKeycode(display, XK_Shift_L), False, 0);
	XSync(display, False);
}
*/
import "C"
import (
	"strconv"
	"time"
	"unsafe"
)

var KeyboardInputUnicodeSleep = time.Millisecond

func keyboardInput(vk uint16, isPress C.int) {
	display := getDisplay()
	C.XTestFakeKeyEvent(display, C.uint(C.XKeysymToKeycode(display, C.ulong(vk))), isPress, 0)
	C.XSync(display, 0)
}

func KeyboardDown(vk uint16) {
	keyboardInput(vk, 1)
	onKeyboardDown(vk)
}

func keyboardUp(vk uint16) {
	keyboardInput(vk, 0)
}

func KeyboardInputString(v string) {
	display := getDisplay()
	for _, r := range v {
		u := C.CString("U" + strconv.FormatInt(int64(r), 16))
		C.keyboardInputUnicode(display, u)
		C.free(unsafe.Pointer(u))
		time.Sleep(KeyboardInputUnicodeSleep)
	}
}
