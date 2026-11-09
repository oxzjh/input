package input

import "sync"

var (
	gMutex        sync.Mutex
	gMouseDown    = map[uint8]struct{}{}
	gKeyboardDown = map[uint16]struct{}{}
)

func onMouseDown(button uint8) {
	gMutex.Lock()
	gMouseDown[button] = struct{}{}
	gMutex.Unlock()
}

func onMouseUp(button uint8) {
	gMutex.Lock()
	delete(gMouseDown, button)
	gMutex.Unlock()
}

func onKeyboardDown(vk uint16) {
	gMutex.Lock()
	gKeyboardDown[vk] = struct{}{}
	gMutex.Unlock()
}

func onKeyboardUp(vk uint16) {
	gMutex.Lock()
	delete(gKeyboardDown, vk)
	gMutex.Unlock()
}

func Release() {
	gMutex.Lock()
	if len(gMouseDown) > 0 {
		for button := range gMouseDown {
			mouseUp(button)
		}
		gMouseDown = map[uint8]struct{}{}
	}
	if len(gKeyboardDown) > 0 {
		for vk := range gKeyboardDown {
			keyboardUp(vk)
		}
		gKeyboardDown = map[uint16]struct{}{}
	}
	gMutex.Unlock()
}
