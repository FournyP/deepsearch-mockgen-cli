package interfaces

type MockPathBuilderInterface interface {
	Build(searchDir, outputDir, interfacePath, interfaceName string) string
}
