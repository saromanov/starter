package models

// ProjectType defines type of the project
type ProjectType int

const (
	Binary ProjectType = iota + 1
	Library
	Unknown
)

// Project provides definition of the project
type Project struct {
	Name        string
	Author      string
	GitPath     string
	Description string
	Dockerfile  string
	Makefile    string
	Type        ProjectType
	SubDirs     []string
	EntryFile   string
}

// StrToProjectType provides converting of
// input string to project type
func StrToProjectType(s string) ProjectType {
	switch s {
	case "binary":
		return Binary
	case "library":
		return Library
	}
	return Unknown
}
