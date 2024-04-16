package slice

func Chunk[T comparable](dataList []T, size int) [][]T {
	chunkCount := len(dataList)/size + 1
	newDataList := make([][]T, 0)

	for i := 0; i < chunkCount; i++ {
		row := make([]T, 0)
		for j := 0; j < size; j++ {
			dataListIndex := i*size + j
			if dataListIndex < len(dataList) {
				row = append(row, dataList[dataListIndex])
			}
		}
		newDataList = append(newDataList, row)
	}

	return newDataList
}
