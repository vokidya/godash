package slice

func Map[T any, R any](dataList []T, iterator func(_index int, _item T) R) []R {
	newDataList := make([]R, 0)
	for index, item := range dataList {
		newDataList = append(newDataList, iterator(index, item))
	}

	return newDataList
}
