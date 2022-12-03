package main

import (
	"fmt"
	//"strings"
)

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(isPalindrome("anna"))
	//fmt.Println(isPalindrome("reconocer"))
	// //////////////////////////////////////////////////////////////////
	//Inicio()
	//fmt.Println(strings.Contains("mi casa verde esta alli", "de"))
	//fmt.Println(strings.Contains("mi casa", "ca"))

	//Repetidas()
	myWord := "G"
	myPhrase := "Atenas dios Griega"

	fmt.Println(Existchar(myPhrase, myWord))

	palbra := "quiero"
	frase := "quiero comida rapida"
	fmt.Println(ExistWord(frase, palbra))

}

func Inicio() {
	fmt.Println("***************************************")
	fmt.Println("Programa que valida una palabra dentro de una frase")
	fmt.Println("Retornando un valor BOOLEANO")
	fmt.Println("*************************************************** ")

	///////
	fmt.Println("***************************************")
	fmt.Println("Array extrae letras repetidos contiguas")
	fmt.Println("***************************************")
}

func Repetidas() {
	//var j, cont int
	fmt.Println(" Array inicial supuesta data repetida ...  R1")
	ar1 := []byte{'a', 'a', 'a', 'b', 'c', 'd', 'd', 'c'}
	//R2 := []byte{' '}
	//Res := []byte{'x'}
	chars := []byte(ar1)
	//dup := []byte(R2)
	//
	fmt.Println(chars)
	fmt.Println("Array pa")
	var a [2]string
	a[0] = "hellow"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	for _, v := range a {
		if v == a[0] {
			fmt.Println("estoy en el array", a[0])
		}
	}

}

func Existchar(cadena string, letra string) bool {

	for i := 0; i < len(cadena); i++ {
		i++
		//fmt.Println("recorrido dentro del FOR cadena", string(cadena[i-1]))
		if string(cadena[i-1]) == string(letra) {
			//fmt.Println("estoy en el array", string(v))
			return true
		}
	}
	return false
}

func ExistWord(frase string, palabra string) bool {
	/// palbra = "rapida"
	/// frase= "quiero comida rapida"
	contador := 0
	//fmt.Println("cantida de palabras ", len(palabra))

	for i := 0; i < len(palabra); i++ {

		for j := 0; j < len(frase); j++ {
			//fmt.Println("recorrido dentro del FOR j", frase[j])
			if string(palabra[i]) == string(frase[j]) && frase[j] != 97 {
				//fmt.Println("recorrido dentro del FOR contador ", contador)
				//fmt.Println("recorrido dentro del FOR i", string(palabra[i]))
				//fmt.Println("recorrido dentro del FOR j", frase[j])
				contador++
			}

			if len(palabra) == contador {
				return true
			}

		}

	}

	return false
}
