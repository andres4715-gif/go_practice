package practica

import "errors"

func RetornandoArray(dataSlice []int) ([]int, error) {
	if len(dataSlice) == 0  {
		return nil, errors.New("💥💥💥 The slice is empty")
	}

	newData := append(dataSlice, 6, 7, 8, 9, 10)
	dataSlice = newData
	return dataSlice, nil
}
