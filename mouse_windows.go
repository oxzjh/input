package input

import (
	"unsafe"
)

var (
	gUser32        *user32
	screenX        int32
	screenY        int32
	mouseInputSize = int32(unsafe.Sizeof(MOUSE_INPUT{}))
)

func init() {
	var err error
	if gUser32, err = newUser32(); err != nil {
		panic(err)
	}
	ResetScreen()
}

func ResetScreen() {
	screenX = gUser32.GetSystemMetrics(SM_CXSCREEN)
	screenY = gUser32.GetSystemMetrics(SM_CYSCREEN)
}

func mouseInput(dx, dy int32, mouseData uint32, flags uint32) {
	gUser32.SendInput(1, unsafe.Pointer(&MOUSE_INPUT{
		Type: INPUT_MOUSE,
		Mi: MOUSEINPUT{
			Dx:        dx,
			Dy:        dy,
			MouseData: mouseData,
			DwFlags:   flags,
		},
	}), mouseInputSize)
}

func mouseMove(dx, dy int32) {
	mouseInput(dx, dy, 0, MOUSEEVENTF_ABSOLUTE|MOUSEEVENTF_MOVE|MOUSEEVENTF_VIRTUALDESK)
}

func MouseMove(x, y int32) {
	mouseMove(65536*x/screenX, 65536*y/screenY)
}

func MouseMoveUint16(x, y uint16) {
	mouseMove(int32(x), int32(y))
}

func MouseMoveWithRatio(x, y float32) {
	mouseMove(int32(65536*x), int32(65536*y))
}

func MouseDown(button uint8) {
	switch button {
	case 0:
		mouseInput(0, 0, 0, MOUSEEVENTF_LEFTDOWN)
	case 1:
		mouseInput(0, 0, 0, MOUSEEVENTF_MIDDLEDOWN)
	case 2:
		mouseInput(0, 0, 0, MOUSEEVENTF_RIGHTDOWN)
	}
	onMouseDown(button)
}

func mouseUp(button uint8) {
	switch button {
	case 0:
		mouseInput(0, 0, 0, MOUSEEVENTF_LEFTUP)
	case 1:
		mouseInput(0, 0, 0, MOUSEEVENTF_MIDDLEUP)
	case 2:
		mouseInput(0, 0, 0, MOUSEEVENTF_RIGHTUP)
	}
}

func MouseWheel(delta int32) {
	mouseInput(0, 0, uint32(delta), MOUSEEVENTF_WHEEL)
}
