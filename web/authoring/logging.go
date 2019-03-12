package authoring

import (
	"log"
	"time"

	"github.com/pamelag/go-web/web/rde"
)

type loggingService struct {
	next Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(s Service) Service {
	return &loggingService{s}
}

func (s *loggingService) AddProject(name string, desc string, createdOn time.Time) (pID rde.ProjectID, err error) {
	defer func(begin time.Time) {
		log.Println("method", "AddProject", "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.next.AddProject(name, desc, createdOn)
}

func (s *loggingService) AddFeature(pID rde.ProjectID, title string) (fID rde.FeatureID, err error) {
	defer func(begin time.Time) {
		log.Println("method", "AddFeature", "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.next.AddFeature(pID, title)
}

func (s *loggingService) AddWireframe(pID rde.ProjectID, fID rde.FeatureID, t string) (wID rde.WireframeID, err error) {
	defer func(begin time.Time) {
		log.Println("method", "AddWireframe", "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.next.AddWireframe(pID, fID, t)
}

func (s *loggingService) UpdateWireframeTitle(pID rde.ProjectID, fID rde.FeatureID, wID rde.WireframeID, t string) (err error) {
	defer func(begin time.Time) {
		log.Println("method", "UpdateWireframeTitle", "took", time.Since(begin), "err", err)
	}(time.Now())
	return s.next.UpdateWireframeTitle(pID, fID, wID, t)
}
