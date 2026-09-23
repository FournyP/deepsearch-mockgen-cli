package interfaces

type InterfaceFinderInterface interface {
	// Find returns every interface declared under root, keyed by name, with its source file path.
	Find(root string) (map[string]string, error)
}
