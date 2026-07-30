package chains

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ComparingChains() {
	fmt.Println(strings.Compare("a", "a")) // When the compared strings are equal the output is: 0
	fmt.Println(strings.Compare("a", "b")) // When the compared strings are not equal the output is: -1
}
func SubstringCount() {
	data := "run"
	dataToFind := strings.Count("run the example to run the test suite", data)
	fmt.Println("Word: ", data, "is present", dataToFind, "times on the provided sentences")
}

func CountSubStringWithSpecificLetter() {
	word := "a"
	textChain := "have a good day"

	characters := utf8.RuneCountInString(textChain)
	fmt.Println("The provided text has", characters, "characters")
	data := strings.Count(textChain, word)

	if data > 1 {
		fmt.Println("Word:", word, ", is present", data, " times in the provided text")
	} else {
		fmt.Println("Word:", word, ", is present", data, " times in the provided text")
	}
}

func ContainsWord() {
	checkWordExists := strings.Contains("if you laugh you have fun", "laugh")
	fmt.Println("The sentence contains a specific word:", checkWordExists)
}

func WordPosition() {
	position := strings.Index("the crazy life", "life")
	fmt.Println("The position is:", position)
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
