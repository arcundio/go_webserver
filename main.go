package main

import (
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PageData struct {
	HostName string
	Images   []ImagenBase64
}

type ImagenBase64 struct {
	Encoding template.URL
	Nombre string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

}

func main() {
	carpeta := os.Args[1]
	puerto := os.Args[2]

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

	var listaGenerada []ImagenBase64

	for i := 0; i < 4; i++ {
		var imagen_aleatoria = archivos[rand.IntN(len(archivos)-1)]

		f, err := os.ReadFile(carpeta + imagen_aleatoria)
		check(err)

		var src = "data:image/jpg;base64," + b64.StdEncoding.EncodeToString(f)

		imagen := ImagenBase64{
			Encoding: template.URL(src),
			Nombre: imagen_aleatoria,
		}

		listaGenerada = append(listaGenerada, imagen)
	}

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			HostName: nombreHost,
			Images:   listaGenerada,
		}

		tmpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":"+puerto, nil))

}
