package slice

func Pick[T map[string]any](sourceObj T, keyList []string) T {
	newObj := map[string]any{}

	for k, v := range sourceObj {
		if Includes(keyList, k) {
			newObj[k] = v
		}
	}

	return newObj
}
