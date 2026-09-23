package settings

type GenerationSettings struct {
	SearchDir string
	OutputDir string
}

func NewGenerationSettings(searchDir, outputDir string) *GenerationSettings {
	return &GenerationSettings{SearchDir: searchDir, OutputDir: outputDir}
}
