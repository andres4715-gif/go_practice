package op_mat // <- mismo paquete que methods.go

import "testing" // <- de aquí sale todo lo necesario para testear

// Un test básico: nombre empieza con "Test" y recibe *testing.T
func TestAdd(t *testing.T) {
	resultado := Add(2, 3)
	esperado := 5

	if resultado != esperado {
		// t es cómo reportamos el fallo (marca fallido pero sigue)
		t.Errorf("Add(2, 3) = %d; se esperaba %d", resultado, esperado)
	}
}

func TestSubtraction(t *testing.T) {
	resultado := Subtraction(10, 4)
	esperado := 6

	if resultado != esperado {
		t.Errorf("Subtraction(10, 4) = %d; se esperaba %d", resultado, esperado)
	}
}

func TestTimes(t *testing.T) {
	resultado := Times(3, 4)
	esperado := 12

	if resultado != esperado {
		t.Errorf("Times(3, 4) = %d; se esperaba %d", resultado, esperado)
	}
}

func TestDividedBy(t *testing.T) {
	resultado := DividedBy(20, 5)
	esperado := 4

	if resultado != esperado {
		t.Errorf("DividedBy(20, 5) = %d; se esperaba %d", resultado, esperado)
	}
}

// Este es un "table-driven test": probamos varios casos con un solo Test.
// Es el patrón idiomático en Go cuando quieres cubrir muchos escenarios.
func TestTotalAddition(t *testing.T) {
	casos := []struct {
		nombre   string
		numeros  []int
		esperado int
	}{
		{"sin numeros", []int{}, 0},
		{"un numero", []int{7}, 7},
		{"varios numeros", []int{1, 2, 3, 4}, 10},
	}

	for _, c := range casos {
		// t.Run crea un sub-test con nombre propio (mejor detalle al fallar)
		t.Run(c.nombre, func(t *testing.T) {
			resultado := TotalAddition(c.numeros...)
			if resultado != c.esperado {
				t.Errorf("TotalAddition(%v) = %d; se esperaba %d", c.numeros, resultado, c.esperado)
			}
		})
	}
}
