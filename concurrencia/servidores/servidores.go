package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		//fmt.Println(servidor, "No está disponible")
		canal <- servidor + " No está disponible"
		return
	}

	//fmt.Println(servidor, "Esta funcionando")
	canal <- servidor + " Esta funcionando"
}

func main() {
	inicio := time.Now()

	canales := make(chan string)

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	for _, server := range servidores {
		// sin concurrencia
		// revisarServidor(server)

		// con concurrencia
		go revisarServidor(server, canales)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canales)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución ", tiempoPaso)
}
