package slice

func Uniq[T comparable](dataList []T) []T {
	newDataList := make([]T, 0)

	for _, item := range dataList {
		if !Includes(newDataList, item) {
			newDataList = append(newDataList, item)
		}
	}

	return newDataList
}
