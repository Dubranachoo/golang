vamos a aprender a definir funciones basicas, que retorne que se ejecute. esto puede servir muchisimo

package main

//esta es una fucnion creada por nosotros
func saludo(){
fmt.println("hola bro")
}

//esta funcion es la principal
func main(){

//debemos llamarlas dentro de la funcion principal para que funcione
saludo()

}

para que quede mejor, vamos a darle un dato a la funcion.

package main

//esta es una fucnion creada por nosotros
//dentro del parentesis de la funcion, pondremos el dato y el tipo de dato que se recibira
func saludo(nombre string){
fmt.println("hola bro")
}

//esta funcion es la principal
func main(){

//debemos llamarlas dentro de la funcion principal para que funcione
saludo()

}


44:19
