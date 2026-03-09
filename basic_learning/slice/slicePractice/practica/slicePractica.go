package practica

import "errors"

func RetornandoArray(data []int) ([]int, error) {
	if len(data) == 0  {
		return nil, errors.New("💥💥💥 The slice is empty")
	}

	newData := append(data, 6, 7, 8, 9, 10)
	data = newData
	return data, nil
}
