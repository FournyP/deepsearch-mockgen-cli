package interfaces

type MockgenReporterInterface interface {
	NoInterfacesFound()
	ProgressFailed(err error)
}
