package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Serv   *http.Server
	Logger *log.Logger
}

func (obj *Server) StartServer() error {
	defer obj.WriteLogf("Server started on %s\n", obj.Serv.Addr)
	return obj.Serv.ListenAndServe()
}

func (obj *Server) WriteLogf(sample string, values ...any) {
	obj.Logger.Printf(sample, values)
}

func (obj *Server) WriteLogln(values ...any) {
	obj.Logger.Println(values)
}

func CreateServer(logger *log.Logger) *Server {

	router := http.NewServeMux()
	router.HandleFunc("/", handlers.HandleRoot)
	router.HandleFunc("/upload", handlers.HandleUpload)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{httpServer, logger}

}
