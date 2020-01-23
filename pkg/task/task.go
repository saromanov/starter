// Package task implements tasks for build stage
package task

// Task defines puppet for tasks
type Task interface {
	Do() error
	String() string
}
