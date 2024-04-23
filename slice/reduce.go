package slice

func Reduce[T any, R any](dataList []T, iterator func(previousState R, _index int, _item T) R, initState R) R {
	currentState := initState

	for index, item := range dataList {
		currentState = iterator(currentState, index, item)
	}

	return currentState
}
