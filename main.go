package main

import (
	//b64 "encoding/base64"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	carpeta := os.Args[1]

	directorio, err := os.Open(carpeta)
	check(err)
	defer directorio.Close()
	nombres, err := directorio.Readdirnames(0)
	check(err)

	var archivos []string
	for _, nombre := range nombres {

		if strings.HasSuffix(nombre, ".jpg") ||
			strings.HasSuffix(nombre, ".png") ||
			strings.HasSuffix(nombre, ".jpeg") {
			archivos = append(archivos, nombre)
		}

	}

	fmt.Println("Cantidad de archivos en la carpeta: ", len(archivos))

	var imagen_aleatoria = archivos[rand.IntN(len(archivos)-1)]

	fmt.Println(imagen_aleatoria)

	nombreHost, err := os.Hostname()
	check(err)

	fmt.Println("Nombre del host: ", nombreHost)

	/*
	f, err := os.ReadFile(carpeta + imagen_aleatoria)
	check(err)

	sEnc := b64.StdEncoding.EncodeToString([]byte(f))
	fmt.Print(sEnc)
	*/
}
