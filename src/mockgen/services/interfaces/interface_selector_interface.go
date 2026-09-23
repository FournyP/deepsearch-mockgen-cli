package interfaces

type InterfaceSelectorInterface interface {
	// Select lets the user pick interfaces among name -> source path and returns the picked names.
	Select(interfaces map[string]string) ([]string, error)
}
