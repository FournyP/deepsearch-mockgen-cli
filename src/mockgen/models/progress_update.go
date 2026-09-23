package models

// ProgressUpdate is sent for each finished mock generation.
type ProgressUpdate struct {
	Name string
	Err  error
}
