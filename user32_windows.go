package input

import (
	"syscall"
	"unsafe"
)

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1

	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1

	MOUSEEVENTF_ABSOLUTE    = 0x8000
	MOUSEEVENTF_MOVE        = 0x0001
	MOUSEEVENTF_LEFTDOWN    = 0x0002
	MOUSEEVENTF_LEFTUP      = 0x0004
	MOUSEEVENTF_RIGHTDOWN   = 0x0008
	MOUSEEVENTF_RIGHTUP     = 0x0010
	MOUSEEVENTF_MIDDLEDOWN  = 0x0020
	MOUSEEVENTF_MIDDLEUP    = 0x0040
	MOUSEEVENTF_VIRTUALDESK = 0x4000
	MOUSEEVENTF_WHEEL       = 0x0800

	KEYEVENTF_KEYUP   = 0x0002
	KEYEVENTF_UNICODE = 0x0004
)

type MOUSE_INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
}

type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type KEYBD_INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
	Unused      [8]byte
}

type user32 struct {
	dll              *syscall.DLL
	getSystemMetrics *syscall.Proc
	sendInput        *syscall.Proc
}

func (u *user32) GetSystemMetrics(nIndex uintptr) int32 {
	ret, _, _ := syscall.SyscallN(u.getSystemMetrics.Addr(), nIndex)
	return int32(ret)
}

func (u *user32) SendInput(nInputs uint32, pInputs unsafe.Pointer, cbSize int32) uint32 {
	ret, _, _ := syscall.SyscallN(u.sendInput.Addr(), uintptr(nInputs), uintptr(pInputs), uintptr(cbSize))
	return uint32(ret)
}

func (u *user32) Close() error {
	return u.dll.Release()
}

func newUser32() (*user32, error) {
	dll, err := syscall.LoadDLL("user32.dll")
	if err != nil {
		return nil, err
	}
	return &user32{
		dll:              dll,
		getSystemMetrics: dll.MustFindProc("GetSystemMetrics"),
		sendInput:        dll.MustFindProc("SendInput"),
	}, nil
}
