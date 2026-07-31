package sliceInitialPractice

import "errors"

func ReturningSliceWithData(dataSlice []int) ([]int, error) {
	if len(dataSlice) == 0 {
		return nil, errors.New("💥💥💥 The slice is empty")
	}

	newData := append(dataSlice, 6, 7, 8, 9, 10)
	return newData, nil
}

func ReturningSliceWithDeleteData(dataSlice []int) ([]int, error) {
	if len(dataSlice) == 0 {
		return nil, errors.New("💥💥💥 The slice is empty")
	}

	newData1 := append(dataSlice, 6, 7, 8, 9, 10)
	dataSlice1 := append(newData1[:1], newData1[2:]...)
	return dataSlice1, nil
}
