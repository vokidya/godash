package slice

func Filter[T any](dataList []T, predicate func(_index int, _item T) bool) []T {
	newDataList := make([]T, 0)

	for index, item := range dataList {
		if predicate(index, item) {
			newDataList = append(newDataList, item)
		}
	}

	return newDataList
}
