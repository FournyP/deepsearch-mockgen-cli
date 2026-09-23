package models

// MockTarget is one interface to mock and where its mock is written.
type MockTarget struct {
	InterfaceName string
	SourcePath    string
	MockPath      string
}
