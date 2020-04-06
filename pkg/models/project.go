package models

// ProjectType defines type of the project
type ProjectType int

const (
	Binary ProjectType = iota + 1
	Library
	Unknown
)

func (p ProjectType) String() string {
	switch p {
	case Binary:
		return "binary"
	case Library:
		return "library"
	}
	return "unknown"
}

// Project provides definition of the project
type Project struct {
	Name               string
	Author             string
	GitPath            string
	Description        string
	Dockerfile         string
	Makefile           string
	Type               ProjectType
	SubDirs            []string
	EntryFile          string
	Commands           []Command
	HostingDescription string
	Badges             []string
	Username           string
}

type Command struct {
	Name string
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
	return Library
}
