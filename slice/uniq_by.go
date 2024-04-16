package slice

func UniqBy[T comparable](dataList []T, isEqual func(item1 T, item2 T) bool) []T {
	newDataList := make([]T, 0)

	for _, dataItem := range dataList {
		if !IncludesBy(newDataList, func(_ int, newDataItem T) bool {
			return isEqual(newDataItem, dataItem)
		}) {
			newDataList = append(newDataList, dataItem)
		}
	}

	return newDataList
}
