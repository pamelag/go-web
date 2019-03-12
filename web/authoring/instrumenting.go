package authoring

import (
	"expvar"
	"time"

	"github.com/pamelag/go-web/web/rde"
)

type instrumentingService struct {
	addProject           *expvar.Int
	addFeature           *expvar.Int
	addWireframe         *expvar.Int
	updateWireframeTitle *expvar.Int
	next                 Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(ap *expvar.Int, af *expvar.Int, aw *expvar.Int, uw *expvar.Int, s Service) Service {
	return &instrumentingService{
		addProject:           ap,
		addFeature:           af,
		addWireframe:         aw,
		updateWireframeTitle: uw,
		next:                 s,
	}
}

func (s *instrumentingService) AddProject(name string, desc string, createdOn time.Time) (pID rde.ProjectID, err error) {
	defer func() {
		s.addProject.Add(1)
	}()
	return s.next.AddProject(name, desc, createdOn)
}

func (s *instrumentingService) AddFeature(pID rde.ProjectID, title string) (fID rde.FeatureID, err error) {
	defer func() {
		s.addFeature.Add(1)
	}()
	return s.next.AddFeature(pID, title)
}

func (s *instrumentingService) AddWireframe(pID rde.ProjectID, fID rde.FeatureID, t string) (wID rde.WireframeID, err error) {
	defer func() {
		s.addWireframe.Add(1)
	}()
	return s.next.AddWireframe(pID, fID, t)
}

func (s *instrumentingService) UpdateWireframeTitle(pID rde.ProjectID, fID rde.FeatureID, wID rde.WireframeID, t string) (err error) {
	defer func() {
		s.updateWireframeTitle.Add(1)
	}()
	return s.next.UpdateWireframeTitle(pID, fID, wID, t)
}
