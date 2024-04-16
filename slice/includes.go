package slice

func IncludesBy[T any](dataList []T, predicate func(_index int, _item T) bool) bool {
	for index, item := range dataList {
		if predicate(index, item) {
			return true
		}
	}

	return false
}

func Includes[T comparable](dataList []T, item T) bool {
	for _, dataItem := range dataList {
		if dataItem == item {
			return true
		}
	}

	return false
}
