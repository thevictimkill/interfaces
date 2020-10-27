package main
import "./medios"
import "fmt"

func main(){
	var op string
	flag := 0
	contenido := medios.ContenidoWeb{
		Multimedios: []medios.Multimedia{
		  },
	  }

	

	for flag < 1 { 
		fmt.Println("1.-Agregar Imagen\n2.-Agregar Audio \n3.-Agregar video\n4.-Mostar\n5.-Salir\nOpciones: ")
		fmt.Scanln(&op)
		switch op {
		case "1":
			var titulo,formato,canales string
			fmt.Println("Titulo: ")
			fmt.Scanln(&titulo)
			fmt.Println("Formato: ")
			fmt.Scanln(&formato)
			fmt.Println("Canales: ")
			fmt.Scanln(&canales)

			i01 := medios.Imagen{titulo,formato,canales}
			contenido.Add(&i01)

			
		case "2":
			var titulo,formato string
			var duracion int64
			fmt.Println("Titulo: ")
			fmt.Scanln(&titulo)
			fmt.Println("Formato: ")
			fmt.Scanln(&formato)
			fmt.Println("Duracion: ")
			fmt.Scanln(&duracion)

			a01 := medios.Audio{titulo,formato,duracion}
			contenido.Add(&a01)
		case "3":
			var titulo,formato string
			var frames int64
			fmt.Println("Titulo: ")
			fmt.Scanln(&titulo)
			fmt.Println("Formato: ")
			fmt.Scanln(&formato)
			fmt.Println("Frames: ")
			fmt.Scanln(&frames)

			v01 := medios.Video{titulo,formato,frames}
			contenido.Add(&v01)
		case "4":
			contenido.Mostrar()
		case "5":
			flag=1
		default:
			fmt.Println("Opcion no exise")
			
		}
	} 
	

}