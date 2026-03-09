package cadenas

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ComparandoCadenas() {
	fmt.Println(strings.Compare("a", "a"))
}
func ConteoSubCadenas() {
	data := "texto"
	dato := strings.Count("cadena de texto", data)
	fmt.Println("El dato: ", data, "se encuentra", dato, "en el texto suministrado")
}

func ConteoSubCadenasConLetra() {
	letra := "a"
	cadenaDeTexto := "andres"

	caracteres := utf8.RuneCountInString(cadenaDeTexto)
	fmt.Println("El texto suministrado tiene", caracteres, "caracteres")
	dato := strings.Count(cadenaDeTexto, letra)

	if dato > 1 {
		fmt.Println("La letra:", letra, ", se encuentra", dato, " veces en el texto suministrado")
	} else {
		fmt.Println("La letra:", letra, ", se encuentra", dato, " vez en el texto suministrado")
	}
}

func ContienePalabra() {
	checkWordExists := strings.Contains("la vida loca", "muerte")
	fmt.Println("La palabra existe en el texto:", checkWordExists)
}

func DondeEstaLaPalabra() {
	posicion := strings.Index("la vida loca", "vida")
	fmt.Println("La palabra esta en la posicion:", posicion)
}

func DondeEstaLaUltimaPalabraDeUnaCadena() {
	posicion := strings.LastIndex("la vida loca es una loca vida", "vida")
	fmt.Println("La ultima palabra en la busqueda esta en la posicion:", posicion)
}

func ReemplazarUnaSubCadena() {
	newTexto := strings.Replace("la vida llega y la mueste se va ", "vida", "suegra", 1)
	fmt.Println("La nueva frase es:", newTexto)
}

func ReemplazarTodosLosCaracteresDeUnaSubCadena() {
	newTexto := strings.Replace("hola como estas, yo estoy mejor", "o", "0", -1) // el -1 es para que reemplace todas las ocurrencias
	fmt.Println("La nueva frase es:", newTexto)
}

func HaciendoUnSplit() {
	newTexto := strings.Split("a, b, c", ", ") // El segundo parametro es el identificador para hacer la separacion
	fmt.Println("Se convierte en un array y este es:", newTexto)
}

func MayusculasAndMinusculas() {
	newTextoMayusculas := strings.ToUpper("hola")
	newTextoMinusculas := strings.ToLower("HOLA")

	caser := cases.Title(language.Spanish) // configurar para usar el idioma español en mayúsculas

	newCapital := caser.String("andres rios")
	fmt.Println("Como queda el texto en mayusculas", newTextoMayusculas)
	fmt.Println("Como queda el texto en minusculas", newTextoMinusculas)
	fmt.Println("Como queda el texto con capital", newCapital)
}

func HacerUnTrimParaEspacios() {
	texto1 := "        Lo que me digas esta bien.           "
	texto2 := "Lo que me digas esta bien > ***************" // aqui es para determinar los caracteres a eliminar
	newTextoUsinTrim := strings.TrimSpace(texto1)
	fmt.Println("El texto sin los espacios queda asi: ", newTextoUsinTrim)
	newTextoUsinTrimDerecho := strings.TrimRight(texto2, "*")
	fmt.Println("El texto sin los espacios a la derecha queda asi: ", newTextoUsinTrimDerecho)
}

func ConcatenarCadenas() {
	a := "andres"
	b := " "
	c := "rios"

	newText := a + b + c

	fmt.Println("El texto concatenado queda asi: ", newText)
}
