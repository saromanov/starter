// Package hosting defines remote hosting for upload/create projects
package hosting

// Hosting defines interface for code hosting
type Hosting interface {
	CreateRepo() error
}
