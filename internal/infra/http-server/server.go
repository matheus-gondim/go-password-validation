package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

type HttpServer struct {
	mux    *chi.Mux
	port   string
	logger zerolog.Logger
}

func New(port string, routes []*chi.Mux) *HttpServer {
	r := chi.NewRouter()
	s := &HttpServer{
		mux:  r,
		port: port,
	}

	s.loggerConfig().
		setMiddleware().
		setDefaultRoutes()

	for _, v := range routes {
		r.Mount("/", v)
	}

	return s
}

func (s *HttpServer) Run() {
	s.logger.Info().Msgf("running http server on port:%s", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.port), s.mux))
}

func (s *HttpServer) loggerConfig() *HttpServer {
	s.logger = httplog.NewLogger("password-validation", httplog.Options{
		JSON:            false,
		LogLevel:        "info",
		TimeFieldFormat: "2006-01-02T15:04:05Z07:00",
	})
	s.mux.Use(httplog.RequestLogger(s.logger))
	return s
}

func (s *HttpServer) setMiddleware() *HttpServer {
	s.mux.Use(
		middleware.AllowContentType("application/json"),
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			AllowedMethods: []string{"POST"},
		}),
		middleware.Timeout(time.Second*60),
	)
	return s
}

func (s *HttpServer) setDefaultRoutes() *HttpServer {
	s.mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	s.mux.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	return s
}
