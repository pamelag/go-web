package server

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pamelag/go-web/web/authoring"
)

const envString = "RP_ENV"

// Server holds the dependencies for a HTTP server.
type Server struct {
	Documenting authoring.Service

	router chi.Router
}

// New returns a new HTTP server.
func New(ds authoring.Service) *Server {
	s := &Server{
		Documenting: ds,
	}

	r := chi.NewRouter()

	r.Use(accessControl)
	r.Use(serveStatic)
	r.Use(middleware.Recoverer)

	r.Route("/authoring", func(r chi.Router) {
		h := authoringHandler{s.Documenting}
		r.Mount("/v1", h.router())
	})

	r.Handle("/debug/vars", http.DefaultServeMux)

	r.Get("/", index)

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func serveStatic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fPath := filepath.Clean(r.URL.Path)
		mainjsRequested, err := regexp.MatchString(`^/main.*.js$`, fPath)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		vendorjsRequested, err := regexp.MatchString(`^/vendor.*.js$`, fPath)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		imagesRequested, err := regexp.MatchString(`^/images/`, fPath)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		if mainjsRequested || vendorjsRequested || imagesRequested {
			fp := filepath.Join("static", fPath)
			http.ServeFile(w, r, fp)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env := os.Getenv(envString)
		if env == "dev" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Add("X-Frame-Options", "DENY")
		w.Header().Add("X-Content-Type-Options", "nosniff")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		w.Header().Add("Content-Security-Policy", "frame-ancestors 'none'")

		if r.Method == "OPTIONS" {
			return
		}

		if env == "prod" {
			w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
			w.Header().Set("Access-Control-Allow-Origin", "https://app.rightprism.com")
		}

		h.ServeHTTP(w, r)
	})
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	/* case shipping.ErrUnknownCargo:
		w.WriteHeader(http.StatusNotFound)
	case tracking.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest) */
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	//lp := filepath.Join("static", "index.html")
	//fp := filepath.Join("static", filepath.Clean(r.URL.Path))

	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}
