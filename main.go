package main

import (
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
	"github.com/alf4ridzi/lapaksiswa-golang/routes"
)

func main() {
	database.InitDatabase()
	footer := path.Join("views", "footer.html")
	header := path.Join("views", "header.html")
	t := template.Must(template.ParseFiles(header, footer))
	t.Execute(os.Stdout, nil)

	// fmt.Println("hello world")
	server := http.NewServeMux()
	routes.MapRoutes(server)
	http.ListenAndServe(":8080", server)

}
