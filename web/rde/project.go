package rde

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

// ProjectID uniquely identifies a particular project.
type ProjectID string

// ProjectStatus represents a project status.
type ProjectStatus string

const (
	// ACTIVE represents a WIP project
	ACTIVE ProjectStatus = "ACTIVE"
	// DELETED represents a deleted project
	DELETED ProjectStatus = "DELETED"
	// ARCHIVED represents a completed project
	ARCHIVED ProjectStatus = "ARCHIVED"
)

// FeatureID uniquely identifies a particular feature.
type FeatureID string

// WireframeID uniquely identifies a particular wireframe.
type WireframeID string

// Project is the central class in the domain model.
type Project struct {
	ProjectID     ProjectID
	Name          string
	Description   string
	CreatedOn     time.Time
	Features      map[string]*Feature
	ProjectStatus ProjectStatus
}

func getID() (string, error) {
	uuidBytes, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	uuidString := uuidBytes.String()
	if uuidString == "" {
		return "", errors.New("error generating id")
	}

	return uuidString, nil
}

// NewProjectID generates a new project ID.
func NewProjectID() (ProjectID, error) {
	uuidString, err := getID()
	if err != nil {
		return "", err
	}

	return ProjectID(uuidString), nil
}

// NewProject creates a new project
func NewProject(id ProjectID, name string, description string, createdOn time.Time) *Project {
	features := make(map[string]*Feature)
	return &Project{
		ProjectID:     id,
		Name:          name,
		Description:   description,
		CreatedOn:     createdOn,
		Features:      features,
		ProjectStatus: ACTIVE,
	}
}

// Feature represents a unit of functionality of a software system that satisfies a requirement.
type Feature struct {
	ProjectID ProjectID
	FeatureID FeatureID
	Title     string
	Content   string
	CreatedOn time.Time
}

// NewFeatureID generates a new feature ID.
func NewFeatureID() (FeatureID, error) {
	uuidString, err := getID()
	if err != nil {
		return "", err
	}

	return FeatureID(uuidString), nil
}

// NewFeature creates a new feature
func NewFeature(projectID ProjectID, id FeatureID, title string) *Feature {
	return &Feature{
		ProjectID: projectID,
		FeatureID: id,
		Title:     title,
	}
}

// Wireframe is a page schematic or screen blueprint
type Wireframe struct {
	WireframeID WireframeID
	ProjectID   ProjectID
	FeatureID   FeatureID
	Title       string
	Elements    []Element
}

// NewWireframeID generates a new wireframe ID.
func NewWireframeID() (WireframeID, error) {
	uuidString, err := getID()
	if err != nil {
		return "", err
	}

	return WireframeID(uuidString), nil
}

// NewWireframe creates a new wirefrane
func NewWireframe(projectID ProjectID, featureID FeatureID, id WireframeID, title string) *Wireframe {
	return &Wireframe{
		ProjectID:   projectID,
		FeatureID:   featureID,
		WireframeID: id,
		Title:       title,
	}
}

// Element represents the controls in a wireframe
type Element struct {
	TimestampKey    int
	Type            string
	StrokeStyle     string
	FillStyle       string
	StartX          int
	StartY          int
	Width           int
	Height          int
	Text            string
	Checked         bool
	MinWidth        int
	MinHeight       int
	RoundedCorner   bool
	FontSize        int
	ForegroundColor string
	Resizable       bool
	Border          string
	Editable        bool
}

// ProjectRepository provides access to project store
type ProjectRepository interface {
	Store(project *Project) error
	StoreFeature(feature *Feature) error
	StoreWireframe(wireframe *Wireframe) error
	UpdateWireframe(wireframe *Wireframe) error
	//Find(id ProjectID) (*Project, error)
	//FindAll() []*Project
}
