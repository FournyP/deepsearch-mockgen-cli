package services

import "log"

type MockgenReporter struct{}

func NewMockgenReporter() *MockgenReporter {
	return &MockgenReporter{}
}

func (r *MockgenReporter) NoInterfacesFound() {
	log.Println("No interfaces found")
}

func (r *MockgenReporter) ProgressFailed(err error) {
	log.Printf("Progress UI error: %v", err)
}
