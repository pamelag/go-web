package authoring

import (
	"errors"
	"time"

	"github.com/pamelag/go-web/web/rde"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides project management methods.
type Service interface {
	AddProject(name string, description string, createdOn time.Time) (rde.ProjectID, error)
	AddFeature(projectID rde.ProjectID, title string) (rde.FeatureID, error)
	AddWireframe(projectID rde.ProjectID,
		featureID rde.FeatureID,
		title string) (rde.WireframeID, error)
	UpdateWireframeTitle(projectID rde.ProjectID,
		featureID rde.FeatureID,
		wireframeID rde.WireframeID,
		title string) error
}

type service struct {
	projects rde.ProjectRepository
}

// NewService creates a project service with necessary dependencies.
func NewService(projects rde.ProjectRepository) Service {
	return &service{
		projects: projects,
	}
}

// AddProject is used to create a new project
func (s *service) AddProject(name, description string, createdOn time.Time) (rde.ProjectID, error) {
	id, err := rde.NewProjectID()
	if err != nil {
		return "", err
	}

	project := rde.NewProject(id, name, description, createdOn)
	if err := s.projects.Store(project); err != nil {
		return "", err
	}

	return project.ProjectID, nil
}

// AddFeature is used to create a new feature for a project
func (s *service) AddFeature(projectID rde.ProjectID, title string) (rde.FeatureID, error) {
	id, err := rde.NewFeatureID()
	if err != nil {
		return "", err
	}

	feature := rde.NewFeature(projectID, id, title)
	if err := s.projects.StoreFeature(feature); err != nil {
		return "", err
	}
	return feature.FeatureID, nil
}

// AddWireframe is used to create a new wireframe for a feature in a project
func (s *service) AddWireframe(projectID rde.ProjectID,
	featureID rde.FeatureID,
	title string) (rde.WireframeID, error) {

	id, err := rde.NewWireframeID()
	if err != nil {
		return "", err
	}

	wireframe := rde.NewWireframe(projectID, featureID, id, title)
	if err := s.projects.StoreWireframe(wireframe); err != nil {
		return "", err
	}
	return wireframe.WireframeID, nil
}

// AddWireframe is used to create a new wireframe for a feature in a project
func (s *service) UpdateWireframeTitle(projectID rde.ProjectID,
	featureID rde.FeatureID,
	id rde.WireframeID,
	title string) error {

	wireframe := rde.NewWireframe(projectID, featureID, id, title)
	return s.projects.UpdateWireframe(wireframe)
}
