package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pamelag/go-web/web/authoring"
	"github.com/pamelag/go-web/web/rde"

	"github.com/go-chi/chi"
)

type authoringHandler struct {
	s authoring.Service
	//logger
}

func (h *authoringHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/projects", func(r chi.Router) {
		r.Post("/", h.addProject)
		r.Route("/{projectID}", func(r chi.Router) {

			r.Route("/features", func(r chi.Router) {
				r.Post("/", h.addFeature)
				r.Route("/{featureID}", func(r chi.Router) {

					r.Route("/wireframes", func(r chi.Router) {
						r.Post("/", h.addWireframe)
						r.Put("/{wireframeID}", h.updateWireframeTitle)
					})

					/* r.Route("/images", func(r chi.Router) {
						r.Post("/", h.addImage)
					}) */

				})
			})

			/* r.Route("/diagrams", func(r chi.Router) {
				r.Post("/", h.addDiagram)
			}) */

		})
	})

	return r
}

func (h *authoringHandler) addProject(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request struct {
		Name        string
		Description string
		CreatedOn   time.Time
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		//h.logger.Log("error", err)
		log.Println("error reading project post", err)
		encodeError(ctx, err, w)
		return
	}

	id, err := h.s.AddProject(request.Name, request.Description, request.CreatedOn)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	var response = struct {
		ID rde.ProjectID `json:"project_id"`
	}{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		//h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}

func (h *authoringHandler) addFeature(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request struct {
		ProjectID string
		Title     string
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		//h.logger.Log("error", err)
		log.Println("error reading feature post", err)
		encodeError(ctx, err, w)
		return
	}

	request.ProjectID = chi.URLParam(r, "projectID")
	id, err := h.s.AddFeature(rde.ProjectID(request.ProjectID), request.Title)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	var response = struct {
		ID rde.FeatureID `json:"feature_id"`
	}{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		//h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}

func (h *authoringHandler) addWireframe(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request struct {
		ProjectID string
		FeatureID string
		Title     string
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		//h.logger.Log("error", err)
		log.Println("error reading feature post", err)
		encodeError(ctx, err, w)
		return
	}

	request.ProjectID = chi.URLParam(r, "projectID")
	request.FeatureID = chi.URLParam(r, "featureID")
	id, err := h.s.AddWireframe(rde.ProjectID(request.ProjectID),
		rde.FeatureID(request.FeatureID),
		request.Title)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	var response = struct {
		ID rde.WireframeID `json:"wireframe_id"`
	}{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		//h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}

func (h *authoringHandler) updateWireframeTitle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var request struct {
		ProjectID   string
		FeatureID   string
		WireframeID string
		Title       string
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		//h.logger.Log("error", err)
		log.Println("error reading feature post", err)
		encodeError(ctx, err, w)
		return
	}

	request.ProjectID = chi.URLParam(r, "projectID")
	request.FeatureID = chi.URLParam(r, "featureID")
	request.WireframeID = chi.URLParam(r, "wireframeID")
	err := h.s.UpdateWireframeTitle(rde.ProjectID(request.ProjectID),
		rde.FeatureID(request.FeatureID),
		rde.WireframeID(request.WireframeID),
		request.Title)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	var response = struct {
		Status string `json:"status"`
	}{
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		//h.logger.Log("error", err)
		encodeError(ctx, err, w)
		return
	}
}
