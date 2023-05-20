//Aqui hablaremos de las variables, con ellas podemos almacenar datos de manera volatil (que se puede reemplazar y por otro tipo de dato o entradas) y las constantes las podemos definir pero nadie puede cambiar sus valores

//primero debemos llamar al paquete variables para poder utilizarlas


package main

func main(){
//para definirlas
//debe usarse el tipo de dato que va a almacenar
var variable1 string
var variable2 string

//se declara la variable
variable1 = "nacho"
variable2 = "que onda"

//para simplificar puede usarse := pero solo para declararla
edad := 19

fmt.println(variable1, variable2, edad)

//para definir constantes
const pito := "hola bro"
fmt.println(pito)

}

//todos los elementos que usaremos (como fmt), se importaran solos al codigo.
