package models

// Project provides definition of the project
type Project struct {
	Name       string
	Author     string
	GitPath    string
	Dockerfile bool
	Makefile   bool
}
