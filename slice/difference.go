package slice

func Difference[T comparable](dataList1 []T, dataList2 []T) []T {
	newDataList := make([]T, 0)

	for _, item := range dataList1 {
		if !Includes(dataList2, item) {
			newDataList = append(newDataList, item)
		}
	}

	return newDataList
}
