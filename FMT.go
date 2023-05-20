print: nos imprime en una sola linea (todo junto)

println: se imprimira en otra linea

package main

import "fmt"

func main(){
que tal := "hola xd"
pit := "que tal"

fmt.print(hola)
fmt.print(pit)


nombre := "nacho"
edad := 19

//con el porcentaje, se imprimira in string
//%d se imprimira en entero
//si no sabemos el tipo de dato, se puede usar %v
// \n salto de linea y \t espacio de tabulacion
fmt.printf("hola, %v su edad %v \n", nombre, edad)

//fmt puede guardarse en una variable

mensajito := fmt.printf("hola, %v su edad %v \n", nombre, edad)
fmt.println(mensaje)

//para saber que tipo de variable es, o que tiene almacenada
//esto te dira el tipo de dato que contiene esa variable
fmt.println("nombre: %t \n", nombre)
fmt.println("edad: %t \n", edad)

//tambien podemos hacer entradas por teclado
fmt.print("ingrese otro nombre: ")

//para ello usaremos esto
// con & vamos a poder capturar en esa variable lo que se va a ingresar
fmt.scanln(&nombre)

//y asi se crea un imput
fmt.println("otro nombre: ", nombre) 

}

