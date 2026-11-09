package input

func GetVK(key string) (uint16, bool) {
	vk, ok := keymap[key]
	return vk, ok
}
