package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<html>
		<head>
			<title>UOC</title>
		</head>
		<body>
			<h1>Soy alumno de la UOC</h1>
			<img src="/static/uoc.png" width="300">
		</body>
		</html>
	`)
}

func main() {

	http.HandleFunc("/", handler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Servidor iniciado en puerto 8080")

	http.ListenAndServe(":8080", nil)
}
