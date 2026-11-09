package input

func KeyboardUp(vk uint16) {
	keyboardUp(vk)
	onKeyboardUp(vk)
}
