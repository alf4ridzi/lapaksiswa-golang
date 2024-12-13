package routes

import (
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.Index())
}
