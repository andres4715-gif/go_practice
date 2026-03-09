package practica

import "errors"

func RetornandoArrayConDatos(dataSlice []int) ([]int, error) {
	if len(dataSlice) == 0  {
		return nil, errors.New("💥💥💥 The slice is empty")
	}

	newData := append(dataSlice, 6, 7, 8, 9, 10)
	dataSlice = newData
	return dataSlice, nil
}

func RetornandoArrayConDatosEliminados(dataSlice []int) ([]int , error) {
	if len(dataSlice) == 0 {
		return nil, errors.New("💥💥💥 The slice is empty")
	}

	newData1 := append(dataSlice, 6, 7, 8, 9, 10)
	dataSlice = newData1
	dataSlice1 := append(dataSlice[:1], dataSlice[2:]...)
	
	return dataSlice1, nil
}
