Aqui veremos los operadores de suma, resta, multiplicacion, division, y otros… no hay mucha diferencia en cuanto a otros lenguajes.

package main

func main(){
//declaramos dos variables
a := 20
b := 15

//suma
result := a + b
fmt.println("suma", result)

//resta (como ya definimos la variable, no hace falta que pongamos los : antes del igual
result = a - b
fmt.println("resta", result)

//multiplicacion
result = a * b
fmt.println("mult", result)

//division
result = a / b
fmt.println("division", result)

//division flotante
var result float64 = 5.1 / 2.1
fmt.println("div float", result)

//modulo de una division (resto)
result = 3 % 2
fmt.println("resto de division", result)






}
