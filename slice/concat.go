package slice

func Concat[T comparable](dataList1 []T, dataList2 []T) []T {
	newDataList := append(dataList1, dataList2...)
	return newDataList
}
